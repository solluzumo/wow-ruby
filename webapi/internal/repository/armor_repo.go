package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type ArmorRepository interface {
	GetById(ctx context.Context, id string) (*models.Armor, error)
}
