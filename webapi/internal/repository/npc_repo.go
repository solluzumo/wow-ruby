package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type NpcRepository interface {
	GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[models.Npc], error)
	GetById(ctx context.Context, id string) (*models.Npc, error)
}
