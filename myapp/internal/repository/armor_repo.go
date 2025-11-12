package repository

import (
	"context"
	"wow-ruby/internal/models"
)

type ArmorRepository interface {
	GetById(ctx context.Context, id string) (*models.Armor, error)
}
