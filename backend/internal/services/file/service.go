package services

import (
	"bytes"
	"context"
	"fmt"
	"gridfs-file-manager/internal/models"
	"mime/multipart"
)

type FileRepository interface {
	UploadFile(file *multipart.FileHeader, userId string, isPublic bool) (string, error)
	DownloadOpenFile(ctx context.Context, bucketId string) (*bytes.Buffer, string, error)
	DownloadPrivateFile(ctx context.Context, bucketId string, userId string) (*bytes.Buffer, string, error)
	GetPublicFiles(ctx context.Context) ([]models.File, error)
	GetPrivateFiles(ctx context.Context, userId string) ([]models.File, error)
	GetSharedFiles(ctx context.Context, userId string) ([]models.File, error)
	UpdateSharedFiles(ctx context.Context, userIds []string, sharedFiles []string) error
}

type UserRepository interface {
	GetIdByUsername(ctx context.Context, username string) (string, error)
}

type Deps struct {
	FileRepo FileRepository
	UserRepo UserRepository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) UploadFile(file *multipart.FileHeader, userId string, isPublic bool) (string, error) {
	uploadFile, err := s.FileRepo.UploadFile(file, userId, isPublic)
	if err != nil {
		return "", err
	}
	return uploadFile, nil
}

func (s *Service) DownloadOpenFile(ctx context.Context, bucketId string) (*bytes.Buffer, string, error) {
	file, filename, err := s.FileRepo.DownloadOpenFile(ctx, bucketId)
	if err != nil {
		return nil, "", err
	}
	return file, filename, nil
}

func (s *Service) DownloadPrivateFile(ctx context.Context, bucketId string, userId string) (*bytes.Buffer, string, error) {
	file, filename, err := s.FileRepo.DownloadPrivateFile(ctx, bucketId, userId)
	if err != nil {
		return nil, "", err
	}
	return file, filename, nil
}

func (s *Service) GetPublicFiles(ctx context.Context) ([]models.File, error) {
	files, err := s.FileRepo.GetPublicFiles(ctx)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *Service) GetPrivateFiles(ctx context.Context, userId string) ([]models.File, error) {
	files, err := s.FileRepo.GetPrivateFiles(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf(userId, err)
	}
	return files, nil
}

func (s *Service) UpdateSharedFiles(ctx context.Context, usernames []string, sharedFiles []string) error {
	var ids []string
	for _, username := range usernames {
		id, err := s.UserRepo.GetIdByUsername(ctx, username)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}

	err := s.FileRepo.UpdateSharedFiles(ctx, ids, sharedFiles)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetSharedFiles(ctx context.Context, userId string) ([]models.File, error) {
	files, err := s.FileRepo.GetSharedFiles(ctx, userId)
	if err != nil {
		return nil, err
	}
	return files, nil
}
