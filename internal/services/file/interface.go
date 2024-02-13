package services

import (
	"bytes"
	"context"
	"gridfs-file-manager/internal/models"
	"mime/multipart"
)

type FileManagement interface {
	UploadFile(file *multipart.FileHeader, userId string, isPublic bool) (string, error)
	DownloadOpenFile(ctx context.Context, bucketId string) (*bytes.Buffer, string, error)
	DownloadPrivateFile(ctx context.Context, bucketId string, userId string) (*bytes.Buffer, string, error)
	GetPublicFiles(ctx context.Context) ([]models.File, error)
	GetPrivateFiles(ctx context.Context, userId string) ([]models.File, error)
	GetSharedFiles(ctx context.Context, userId string) ([]models.File, error)
	UpdateSharedFiles(ctx context.Context, usernames []string, sharedFiles []string) error
}
