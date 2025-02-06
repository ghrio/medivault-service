package repository

import (
	"context"
	db "medivault-service/internal/db/generated"
	models "medivault-service/internal/db/generated"
)

// UserRepoImpl implements UserRepository using sqlc
type UserRepoImpl struct {
	Queries *db.Queries // sqlc-generated queries
}

// NewUserRepository creates a new instance of UserRepoImpl
func NewUserRepository(q *db.Queries) UserRepository {
	return &UserRepoImpl{Queries: q}
}

// GetAllUsers fetches all users
func (r *UserRepoImpl) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return r.Queries.GetAllUsers(ctx)
}

// GetActiveUsers fetches only active users
func (r *UserRepoImpl) GetActiveUsers(ctx context.Context) ([]models.User, error) {
	// Fetch from sqlc-generated query
	rows, err := r.Queries.GetActiveUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Convert []db.GetActiveUsersRow to []models.User
	var users []models.User
	for _, row := range rows {
		users = append(users, models.User{
			ID:        row.ID,
			Name:      row.Name,
			Email:     row.Email,
			Password:  row.Password,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			DeletedAt: row.DeletedAt, // Handle NULL if needed
		})
	}

	return users, nil
}

// CreateUser inserts a new user
func (r *UserRepoImpl) CreateUser(ctx context.Context, name, email, password string) (models.User, error) {
	params := db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	}

	user, err := r.Queries.CreateUser(ctx, params)
	if err != nil {
		return models.User{}, err
	}

	// Convert db.User to models.User if needed
	return models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

// UpdateUserEmail updates the email of a user
func (r *UserRepoImpl) UpdateUserEmail(ctx context.Context, id int32, email string) error {
	params := db.UpdateUserEmailParams{
		Email: email,
		ID:    id,
	}
	err := r.Queries.UpdateUserEmail(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

// SoftDeleteUser sets deleted_at to NOW()
func (r *UserRepoImpl) SoftDeleteUser(ctx context.Context, id int32) error {
	return r.Queries.SoftDeleteUser(ctx, id)
}
