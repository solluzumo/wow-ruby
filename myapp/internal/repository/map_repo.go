package repository

import (
	"context"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/models"
)

type MapRepostiory interface {
	GetMapById(ctx context.Context, id string) (*models.Item, error)
	CreateMap(ctx context.Context, req *dto.MapRequest) (*models.Map, error)
}
