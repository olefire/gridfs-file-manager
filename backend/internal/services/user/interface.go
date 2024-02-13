package services

import (
	"context"
	"gridfs-file-manager/internal/models"
)

type UserManagement interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
}
