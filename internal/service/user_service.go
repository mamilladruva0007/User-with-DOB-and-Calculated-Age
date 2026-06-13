package service

import (
	"context"

	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/db/sqlc"
	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUser(ctx context.Context, id int32) (sqlc.User, error) {
	return s.Repo.GetUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]sqlc.User, error) {
	return s.Repo.ListUsers(ctx)
}
