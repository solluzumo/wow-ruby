package repository

import (
	"context"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/models"
)

type NpcRepository interface {
	GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[models.Npc], error)
	GetById(ctx context.Context, id string) (*models.Npc, error)
}
