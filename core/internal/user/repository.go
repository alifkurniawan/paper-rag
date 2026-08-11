package user

import (
	"context"
	"core/internal/db/sqlc"
)

type Repository struct {
	queries *sqlc.Queries
}

func NewRepository(queries *sqlc.Queries) *Repository {
	return &Repository{queries: queries}
}

func toDomainUser(u sqlc.User) User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

func (r *Repository) GetByID(ctx context.Context, id int32) (User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return User{}, err
	}
	return toDomainUser(dbUser), nil
}

func (r *Repository) List(ctx context.Context) ([]User, error) {
	dbUsers, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(dbUsers))
	for _, u := range dbUsers {
		users = append(users, toDomainUser(u))
	}
	return users, nil
}

func (r *Repository) Create(ctx context.Context, name, emil string) (User, error) {
	dbUser, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:  name,
		Email: emil,
	})
	if err != nil {
		return User{}, err
	}

	return toDomainUser(dbUser), nil
}
