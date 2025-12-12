package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/auth/internal/domain"
	"github.com/solluzumo/wow-ruby/auth/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.UserDomain) (string, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateStatus(ctx context.Context, id string) error
}
