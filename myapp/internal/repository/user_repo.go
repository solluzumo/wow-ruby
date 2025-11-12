package repository

import (
	"context"
	"wow-ruby/internal/domain"
	"wow-ruby/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.UserDomain) (string, error)
	FindByLogin(ctx context.Context, login string) (*models.User, error)
	GetById(ctx context.Context, id string) (*models.User, error)
	FindByLoginAndHash(ctx context.Context, login string, hash string) (*models.User, error)
}
