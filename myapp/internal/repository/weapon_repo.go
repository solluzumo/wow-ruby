package repository

import (
	"context"
	"wow-ruby/internal/models"
)

type WeaponRepository interface {
	GetById(ctx context.Context, id string) (*models.Weapon, error)
}
