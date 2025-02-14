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

type PatientCollectionService interface {
	GetPatientCollectionByID(ctx context.Context, id int) (*models.PatientCollection, error)
	GetAllPatientCollections(ctx context.Context) ([]*models.PatientCollection, error)
	CreatePatientCollection(ctx context.Context, patientCollection *models.PatientCollection) error
	UpdatePatientCollection(ctx context.Context, patientCollection *models.PatientCollection) error
	DeletePatientCollection(ctx context.Context, id int) error
}

type RoleService interface {
	GetRoleByID(ctx context.Context, id int) (*models.Role, error)
	GetAllRoles(ctx context.Context) ([]*models.Role, error)
	CreateRole(ctx context.Context, role *models.Role) error
	UpdateRole(ctx context.Context, role *models.Role) error
	DeleteRole(ctx context.Context, id int) error
}
