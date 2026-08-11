package user

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByID(ctx context.Context, id int32) (User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, name, email string) (User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return User{}, errors.New("name is required")
	}

	if email == "" {
		return User{}, errors.New("email is required")
	}

	return s.repo.Create(ctx, name, email)
}
