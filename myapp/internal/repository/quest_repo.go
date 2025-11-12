package repository

import (
	"context"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/models"
)

type QuestRepository interface {
	GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[models.Quest], error)
	GetById(ctx context.Context, id string) (*models.Quest, error)
}
