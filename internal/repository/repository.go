package repository

import (
	"context"
	"medivault-service/internal/db/generated"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]generated.User, error)
	GetActiveUsers(ctx context.Context) ([]generated.User, error)
	GetUserByEmail(ctx context.Context, email string) (generated.User, error)
	GetUserByID(ctx context.Context, id int32) (generated.User, error)
	CreateUser(ctx context.Context, name, email, password string) (generated.User, error)
	UpdateUserEmail(ctx context.Context, id int32, email string) error
	SoftDeleteUser(ctx context.Context, id int32) error
}

type PatientFileRepository interface {
	GetAllPatientFiles(ctx context.Context) ([]generated.PatientFile, error)
	GetPatientFileByID(ctx context.Context, id int) (generated.PatientFile, error)
	CreatePatientFile(ctx context.Context, patientID int, fileID int, description string) (generated.PatientFile, error)
	UpdatePatientFile(ctx context.Context, id int, patientID int, fileID int, description string) (generated.PatientFile, error)
	DeletePatientFile(ctx context.Context, id int) error
}

type PatientCollectionRepository interface {
	GetPatientCollectionByUser(ctx context.Context, userID int32) (generated.PatientCollection, error)
	CreatePatientCollection(ctx context.Context, userID int32) (generated.PatientCollection, error)
}

type RoleRepository interface {
	AssignRoleToUser(ctx context.Context, userID, roleID int32) error
	GetUsersWithRoles(ctx context.Context) ([]generated.GetUsersWithRolesRow, error)
	RemoveRoleFromUser(ctx context.Context, userID, roleID int32) error
}
