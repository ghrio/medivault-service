package repository

import (
	"context"
	models "medivault-service/internal/db/generated"
)

// UserRepository defines the methods for user management
type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetActiveUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, name, email, password string) (models.User, error)
	UpdateUserEmail(ctx context.Context, id int32, email string) error
	SoftDeleteUser(ctx context.Context, id int32) error
}
