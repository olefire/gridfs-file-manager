package services

import (
	"context"
	"gridfs-file-manager/internal/models"
)

type AuthManagement interface {
	SignUpUser(ctx context.Context, signupBody *models.SignupBody) (string, error)
	SignInUser(ctx context.Context, loginBody *models.LoginRequest) (string, error)
}
