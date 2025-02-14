package services

import (
	"context"
	"errors"
	models "medivault-service/internal/db/generated"
	"medivault-service/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService returns a UserService implementation
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return errors.New("not implemented")
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {

	return nil, errors.New("not implemented")
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*models.User, error) {

	return nil, errors.New("not implemented")
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return errors.New("not implemented")
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return errors.New("not implemented")
}
