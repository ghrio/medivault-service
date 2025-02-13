package services

import (
	"context"
	models "medivault-service/internal/db/generated"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}

type PatientFileService interface {
	CreateFile(ctx context.Context, file *models.PatientFile) error
	GetFileByID(ctx context.Context, id int) (*models.PatientFile, error)
	GetAllFiles(ctx context.Context) ([]*models.PatientFile, error)
	UpdateFile(ctx context.Context, file *models.PatientFile) error
}

type PatientCollectionService interface{}
