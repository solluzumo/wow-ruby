package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type QuestRepository interface {
	GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[models.Quest], error)
	GetById(ctx context.Context, id string) (*models.Quest, error)
}
