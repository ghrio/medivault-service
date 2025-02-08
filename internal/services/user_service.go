package services

import (
	"context"
	db "medivault-service/internal/db/generated"
	"medivault-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(ctx context.Context) ([]db.User, error) {
	return s.repo.GetAllUsers(ctx)
}
