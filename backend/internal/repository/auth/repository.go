package repostiry

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gridfs-file-manager/internal/models"
	AuthService "gridfs-file-manager/internal/services/auth"
	"gridfs-file-manager/pkg/utils"
	"time"
)

type AuthRepository struct {
	collection *mongo.Collection
}

var _ AuthService.Repository = (*AuthRepository)(nil)

func NewUserRepository(col *mongo.Collection) *AuthRepository {
	return &AuthRepository{
		collection: col,
	}
}

func (ar *AuthRepository) SignUpUser(ctx context.Context, signupBody *models.SignupBody) (string, error) {
	var user models.User

	if err := ar.collection.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupBody.Email}}).Decode(&user); !errors.Is(err, mongo.ErrNoDocuments) {
		return "", fmt.Errorf("email already exists")
	}

	if err := ar.collection.FindOne(ctx, bson.D{{Key: "username", Value: signupBody.Username}}).Decode(&user); !errors.Is(err, mongo.ErrNoDocuments) {
		return "", fmt.Errorf("username already exists")
	}

	if err := ar.collection.FindOne(ctx, bson.D{{Key: "password", Value: signupBody.Password}}).Decode(user); !errors.Is(err, mongo.ErrNoDocuments) {
		return "", fmt.Errorf("choose another password")
	}

	doc := bson.D{
		{Key: "name", Value: signupBody.Name},
		{Key: "email", Value: signupBody.Email},
		{Key: "username", Value: signupBody.Username},
		{Key: "password", Value: signupBody.Password},
		{Key: "createdAt", Value: time.Now().UTC()},
		{Key: "updatedAt", Value: time.Now().UTC()},
		{Key: "sharedFiles", Value: []primitive.ObjectID{}},
	}

	insertedUser, err := ar.collection.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}
	id := insertedUser.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (ar *AuthRepository) SignInUser(ctx context.Context, loginBody *models.LoginRequest) (string, error) {
	var user models.User
	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: loginBody.Username}},
			bson.D{{Key: "username", Value: loginBody.Username}},
		},
		},
	}
	if err := ar.collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return "", fmt.Errorf("incorrect credentials: %w", err)
	}

	if err := utils.VerifyPassword(user.Password, loginBody.Password); err != nil {
		return "", fmt.Errorf("incorrect credentials: %w", err)
	}

	updateDoc := bson.M{
		"$set": bson.M{
			"isActive": true,
		},
	}

	_, err := ar.collection.UpdateByID(ctx, user.Id, updateDoc)
	if err != nil {
		return "", err
	}

	id := user.Id.Hex()
	return id, nil
}
