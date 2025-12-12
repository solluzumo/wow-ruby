package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type MapRepostiory interface {
	GetMapById(ctx context.Context, id string) (*models.Item, error)
	CreateMap(ctx context.Context, req *dto.MapRequest) (*models.Map, error)
}
