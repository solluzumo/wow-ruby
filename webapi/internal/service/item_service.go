package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/pkg/enums"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/domain"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository"
)

type ItemService struct {
	itemRepo   repository.ItemRepository
	weaponRepo repository.WeaponRepository
	armorRepo  repository.ArmorRepository
}

func NewItemService(itemRepo repository.ItemRepository, weaponRepo repository.WeaponRepository, armorRepo repository.ArmorRepository) *ItemService {
	return &ItemService{
		itemRepo:   itemRepo,
		weaponRepo: weaponRepo,
		armorRepo:  armorRepo,
	}
}

func (is *ItemService) GetItemById(ctx context.Context, id string) (*domain.ItemDomain, error) {

	//Проверяем есть ли предмет и получаем общие данные
	result, err := is.itemRepo.GetItemById(ctx, id)
	if err != nil {
		return nil, err
	}

	//Проверяем тип предмета, если оружие или броня - дополняем ответ данными
	if result.ItemType == string(enums.WeaponType) {
		detail, err := is.weaponRepo.GetById(ctx, id)
		if err != nil {
			return result, err
		}
		result.Weapon = &domain.WeaponDomain{
			Slot:       detail.Slot,
			Durability: detail.Durability,
			Damage:     detail.Damage,
			Speed:      detail.Speed,
			WeaponType: string(detail.WeaponType),
		}

	}

	if result.ItemType == string(enums.ArmorType) {
		detail, err := is.armorRepo.GetById(ctx, id)
		if err != nil {
			return result, err
		}
		result.Armor = &domain.ArmorDomain{
			Slot:       detail.Slot,
			Durability: detail.Durability,
			ArmorType:  string(detail.ArmorType),
			ArmorValue: detail.ArmorValue,
		}

	}

	return result, nil

}

func (is *ItemService) GetItemList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[models.Item], error) {

	item_response, err := is.itemRepo.GetList(ctx, list_request)
	if err != nil {
		return nil, err
	}

	return item_response, nil
}
