package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gridfs-file-manager/internal/models"
	UserService "gridfs-file-manager/internal/services/user"
)

type UserRepository struct {
	collection *mongo.Collection
}

var _ UserService.Repository = (*UserRepository)(nil)

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: col,
	}
}

func (ur *UserRepository) GetUser(ctx context.Context, userId string) (*models.User, error) {
	var user models.User

	idHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	if err := ur.collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&user); err != nil {
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (ur *UserRepository) FindUser(ctx context.Context, field string, value string) (*models.User, error) {
	user := new(models.User)
	filter := bson.M{}
	filter[field] = value
	if err := ur.collection.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, userId string, update bson.M) error {
	idHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	_, err = ur.collection.UpdateByID(ctx, idHex, update)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetSharedFiles(ctx context.Context, userId string) ([]models.File, error) {
	var files []models.File

	idHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": idHex}

	err = ur.collection.FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"sharedFiles": 1, "_id": 1})).Decode(files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (ur *UserRepository) GetIdByUsername(ctx context.Context, username string) (string, error) {
	var body struct {
		Id primitive.ObjectID `bson:"_id" json:"_id"`
	}
	err := ur.collection.FindOne(ctx, bson.M{"username": username}, options.FindOne().SetProjection(bson.M{"_id": 1})).Decode(&body)
	if err != nil {
		return "", fmt.Errorf(body.Id.Hex(), err)
	}
	return body.Id.Hex(), nil
}

//type Body struct {
//	Files []primitive.ObjectID `json:"availableFiles" bson:"availableFiles"`
//}
//
//func (ur *UserRepository) GetSharedFiles(ctx context.Context, userId string) ([]models.File, error) {
//
//	sharedFiles, err := utils.GetSharedFiles(ctx, userId)
//	if err != nil {
//		return err
//	}
//
//	mediaCollection := database.Mi.Db.Collection(database.MediaCollection)
//	var files []bson.M
//	cursor, err := mediaCollection.Aggregate(ctx, []bson.M{
//		{"$match": bson.M{"_id": bson.M{"$in": sharedFiles.Files}}},
//		{"$project": bson.M{"_id": 1, "filename": 1}},
//	})
//	err = cursor.All(ctx, &files)
//	if err != nil {
//		return err
//	}
//
//	return utils.CreatedResponse(c, "Get", fiber.Map{"results": files})
//}
