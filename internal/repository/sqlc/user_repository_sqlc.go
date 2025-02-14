package repository

import (
	"context"
	"medivault-service/internal/db/generated"
	"medivault-service/internal/repository"

	"github.com/jackc/pgx/v5"
)

// UserRepoSqlc implements UserRepository using sqlc
type UserRepoSqlc struct {
	queries *generated.Queries
	conn    *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) repository.UserRepository {
	return &UserRepoSqlc{queries: generated.New(conn), conn: conn}
}

func (u *UserRepoSqlc) GetAllUsers(ctx context.Context) ([]generated.User, error) {
	return nil, nil
}
func (u *UserRepoSqlc) GetActiveUsers(ctx context.Context) ([]generated.User, error) {
	return nil, nil
}
func (u *UserRepoSqlc) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	return generated.User{}, nil
}
func (u *UserRepoSqlc) GetUserByID(ctx context.Context, id int32) (generated.User, error) {
	return generated.User{}, nil
}
func (u *UserRepoSqlc) CreateUser(ctx context.Context, name, email, password string) (generated.User, error) {
	return generated.User{}, nil
}
func (u *UserRepoSqlc) UpdateUserEmail(ctx context.Context, id int32, email string) error {
	return nil
}
func (u *UserRepoSqlc) SoftDeleteUser(ctx context.Context, id int32) error { return nil }
