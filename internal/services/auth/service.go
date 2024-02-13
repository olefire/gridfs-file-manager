package services

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"gridfs-file-manager/internal/models"
	"gridfs-file-manager/pkg/utils"
	"time"
)

type Repository interface {
	SignUpUser(ctx context.Context, signupBody *models.SignupBody) (string, error)
	SignInUser(ctx context.Context, loginBody *models.LoginRequest) (string, error)
}

type Deps struct {
	AuthRepo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) SignUpUser(ctx context.Context, signupBody *models.SignupBody) (string, error) {
	hashPassword, err := utils.HashPassword(signupBody.Password)
	if err != nil {
		return "", err
	}
	signupBody.Password = hashPassword

	id, err := s.AuthRepo.SignUpUser(ctx, signupBody)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) SignInUser(ctx context.Context, loginBody *models.LoginRequest) (string, error) {
	id, err := s.AuthRepo.SignInUser(ctx, loginBody)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	accessToken, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
