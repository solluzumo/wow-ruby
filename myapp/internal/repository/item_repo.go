package repository

import (
	"context"
	"wow-ruby/internal/domain"
)

type ItemRepository interface {
	GetItemById(ctx context.Context, id string) (*domain.ItemDomain, error)
	GetList(ctx context.Context, req *domain.ListDomain) (*domain.ListResponseDomain[domain.ItemDomain], error)
}
