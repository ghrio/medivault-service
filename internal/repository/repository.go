package repository

import (
	"context"
	models "medivault-service/internal/db/generated"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetActiveUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, name, email, password string) (models.User, error)
	UpdateUserEmail(ctx context.Context, id int32, email string) error
	SoftDeleteUser(ctx context.Context, id int32) error
}

type PatientFileRepository interface {
	GetAllPatientFiles(ctx context.Context) ([]models.PatientFile, error)
	GetPatientFileByID(ctx context.Context, id int) (models.PatientFile, error)
	CreatePatientFile(ctx context.Context, patientID int, fileID int, description string) (models.PatientFile, error)
	UpdatePatientFile(ctx context.Context, id int, patientID int, fileID int, description string) (models.PatientFile, error)
	DeletePatientFile(ctx context.Context, id int) error
}

type PatientCollectionRepository interface {
	GetPatientCollectionByUser(ctx context.Context, userID int32) (models.PatientCollection, error)
	CreatePatientCollection(ctx context.Context, userID int32) (models.PatientCollection, error)
}

type RoleRepository interface {
	AssignRoleToUser(ctx context.Context, userID, roleID int32) error
	GetUsersWithRoles(ctx context.Context) ([]models.GetUsersWithRolesRow, error)
	RemoveRoleFromUser(ctx context.Context, userID, roleID int32) error
}
