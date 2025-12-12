package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type WeaponRepository interface {
	GetById(ctx context.Context, id string) (*models.Weapon, error)
}
