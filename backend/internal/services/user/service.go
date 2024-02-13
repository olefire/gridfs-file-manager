package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"gridfs-file-manager/internal/models"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
	FindUser(ctx context.Context, field string, value string) (*models.User, error)
	UpdateUser(ctx context.Context, userId string, update bson.M) error
	GetIdByUsername(ctx context.Context, username string) (string, error)
	//UpdateAvailableFiles(ctx context.Context, users []string, availableFiles []string) error
	//GetSharedFiles(ctx context.Context, userId string) ([]models.File, error)
}

type Deps struct {
	UserRepo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can`t get user: %w", err)
	}
	return user, nil
}

func (s *Service) FindUser(ctx context.Context, field string, value string) (*models.User, error) {
	user, err := s.UserRepo.FindUser(ctx, field, value)
	if err != nil {
		return nil, fmt.Errorf("can`t get user: %w", err)
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, userId string, update bson.M) error {
	err := s.UserRepo.UpdateUser(ctx, userId, update)
	if err != nil {
		return fmt.Errorf("can`t update user: %w", err)
	}
	return nil
}

//func (s *Service) UpdateAvailableFiles(ctx context.Context, users []string, sharedFiles []string) error {
//	err := s.UserRepo.UpdateAvailableFiles(ctx, users, sharedFiles)
//	if err != nil {
//		return fmt.Errorf("can`t update user: %w", err)
//	}
//	return nil
//}

//func (s *Service) GetSharedFiles(ctx context.Context, userId string) ([]models.File, error) {
//	files, err := s.UserRepo.GetSharedFiles(ctx, userId)
//	if err != nil {
//		return nil, fmt.Errorf("can`t get files: %w", err)
//	}
//	return files, nil
//}
