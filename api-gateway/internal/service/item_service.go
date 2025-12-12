package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/gateway/internal/domain"
	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
)

type ItemProvider interface {
}

type ItemService struct {
	item ItemProvider
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (is *ItemService) GetItemById(ctx context.Context, id string) (*domain.ItemDomain, error) {

	return nil, nil

}

func (is *ItemService) GetItemList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[domain.ItemDomain], error) {

	return nil, nil
}
