package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID   `bson:"_id" json:"_id"`
	Name           string               `bson:"name" json:"name"`
	Username       string               `bson:"username" json:"username"`
	Password       string               `bson:"password" json:"password"`
	Email          string               `bson:"email" json:"email"`
	IsActive       bool                 `bson:"isActive" json:"isActive"`
	CreatedAt      time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt" json:"updatedAt"`
	AvailableFiles []primitive.ObjectID `json:"availableFiles" bson:"availableFiles"`
}
