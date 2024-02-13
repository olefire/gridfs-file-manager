package repository

import (
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gridfs-file-manager/internal/models"
	FileService "gridfs-file-manager/internal/services/file"
	"gridfs-file-manager/pkg/utils"
	"io"
	"mime/multipart"
)

type FileRepository struct {
	gridFS       *mongo.Collection
	urCollection *mongo.Collection
	bucket       *gridfs.Bucket
}

var _ FileService.FileRepository = (*FileRepository)(nil)

func NewFileRepository(gridFSCol *mongo.Collection, sfCol *mongo.Collection, bucket *gridfs.Bucket) *FileRepository {
	return &FileRepository{
		gridFS:       gridFSCol,
		urCollection: sfCol,
		bucket:       bucket,
	}
}

func (fr *FileRepository) UploadFile(file *multipart.FileHeader, userId string, isPublic bool) (string, error) {
	filename := file.Filename

	fileData, err := file.Open()
	if err != nil {
		return "", err
	}

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	uploadOptions := options.GridFSUpload().SetMetadata(bson.M{"userId": id, "isPublic": isPublic})

	bucketId, err := fr.bucket.UploadFromStream(filename, io.Reader(fileData), uploadOptions)
	if err != nil {
		return "", err
	}
	return bucketId.Hex(), nil
}

func (fr *FileRepository) DownloadOpenFile(ctx context.Context, bucketId string) (*bytes.Buffer, string, error) {
	id, err := primitive.ObjectIDFromHex(bucketId)
	if err != nil {
		return nil, "", err
	}

	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.M{"filename": 1, "metadata.isPublic": 1})
	var result bson.M

	err = fr.gridFS.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return nil, "", err
	}

	filename := result["filename"].(string)
	isPublic := result["metadata"].(bson.M)["isPublic"]
	if isPublic == "false" {
		return nil, "", fmt.Errorf("this is private file: %w", err)
	}

	var buffer bytes.Buffer
	if _, err := fr.bucket.DownloadToStream(id, &buffer); err != nil {
		return nil, "", err
	}

	return &buffer, filename, nil
}

func (fr *FileRepository) DownloadPrivateFile(ctx context.Context, bucketId string, userId string) (*bytes.Buffer, string, error) {
	id, err := primitive.ObjectIDFromHex(bucketId)
	if err != nil {
		return nil, "", err
	}

	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.M{"filename": 1, "metadata.userId": 1})
	var result bson.M

	err = fr.gridFS.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		return nil, "", err
	}

	filename := result["filename"].(string)
	authorId := result["metadata"].(bson.M)["userId"].(primitive.ObjectID).Hex()

	if authorId != userId {
		return nil, "", fmt.Errorf("you do not have access to this file: %w", err)
	}

	var buffer bytes.Buffer
	if _, err := fr.bucket.DownloadToStream(id, &buffer); err != nil {
		return nil, "", err
	}

	return &buffer, filename, nil
}

func (fr *FileRepository) DownloadSharedFile(ctx context.Context, bucketId string, userId string) (*bytes.Buffer, string, error) {
	uId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, "", err
	}
	bId, err := primitive.ObjectIDFromHex(bucketId)
	if err != nil {
		return nil, "", err
	}

	filter := bson.M{
		"_id":         uId,
		"sharedFiles": bId}
	count, err := fr.urCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, "", err
	}
	if count == 0 {
		return nil, "", fmt.Errorf("no filey")
	}

	filter = bson.M{"_id": bId}
	opts := options.FindOne().SetProjection(bson.M{"filename": 1})
	var result bson.M

	if err = fr.gridFS.FindOne(ctx, filter, opts).Decode(&result); err != nil {
		return nil, "", err
	}

	filename := result["filename"].(string)

	var buffer bytes.Buffer
	if _, err := fr.bucket.DownloadToStream(bId, &buffer); err != nil {
		return nil, "", err
	}

	return &buffer, filename, nil
}

func (fr *FileRepository) GetPrivateFiles(ctx context.Context, userId string) ([]models.File, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"metadata.userId": id}

	cursor, err := fr.gridFS.Find(ctx, filter, options.Find().SetProjection(bson.M{"filename": 1, "_id": 1}))
	if err != nil {
		return nil, err
	}

	var files []models.File
	for cursor.Next(ctx) {
		var file models.File
		err = cursor.Decode(&file)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (fr *FileRepository) GetPublicFiles(ctx context.Context) ([]models.File, error) {
	filter := bson.M{"metadata.isPublic": "true"}

	cursor, err := fr.gridFS.Find(ctx, filter, options.Find().SetProjection(bson.M{"filename": 1, "_id": 1}))
	if err != nil {
		return nil, err
	}

	var results []models.File
	for cursor.Next(ctx) {
		var result models.File
		err = cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (fr *FileRepository) GetSharedFiles(ctx context.Context, userId string) ([]models.File, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	var body struct {
		Files []primitive.ObjectID `json:"sharedFiles" bson:"sharedFiles"`
	}

	err = fr.urCollection.FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"sharedFiles": 1})).Decode(&body)
	if err != nil {
		return nil, err
	}

	var files []models.File
	cursor, err := fr.gridFS.Aggregate(ctx, []bson.M{
		{"$match": bson.M{"_id": bson.M{"$in": body.Files}}},
		{"$project": bson.M{"_id": 1, "filename": 1}},
	})
	err = cursor.All(ctx, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (fr *FileRepository) UpdateSharedFiles(ctx context.Context, userIds []string, sharedFiles []string) error {
	files, err := utils.ConvertToPrimitiveObjectId(sharedFiles)
	if err != nil {
		return err
	}
	pUserIds, err := utils.ConvertToPrimitiveObjectId(userIds)
	if err != nil {
		return err
	}

	for _, id := range pUserIds {
		for i := 0; i < len(files); i++ {
			_, err := fr.urCollection.UpdateOne(
				ctx,
				bson.M{"_id": id},
				bson.M{"$addToSet": bson.M{"sharedFiles": files[i]}},
				options.Update().SetUpsert(true),
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
