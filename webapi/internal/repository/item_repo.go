package repository

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/domain"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type ItemRepository interface {
	GetItemById(ctx context.Context, id string) (*domain.ItemDomain, error)
	GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[models.Item], error)
}
