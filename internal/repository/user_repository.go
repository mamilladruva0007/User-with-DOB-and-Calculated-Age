package repository

import (
	"context"

	"github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age/db/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		Queries: q,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams) error {
	_, err := r.Queries.CreateUser(ctx, params)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, id int32) (sqlc.User, error) {
	return r.Queries.GetUser(ctx, id)
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]sqlc.User, error) {
	return r.Queries.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams) error {
	return r.Queries.UpdateUser(ctx, params)
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.Queries.DeleteUser(ctx, id)
}
