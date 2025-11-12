package service

import (
	"context"
	"wow-ruby/internal/domain"
	"wow-ruby/internal/pkg/enums"
	"wow-ruby/internal/repository"
)

type ItemService struct {
	itemRepo   repository.ItemRepository
	weaponRepo repository.WeaponRepository
	armorRepo  repository.ArmorRepository
	jwt        string
}

func NewItemService(itemRepo repository.ItemRepository, weaponRepo repository.WeaponRepository, armorRepo repository.ArmorRepository, jwt string) *ItemService {
	return &ItemService{
		itemRepo:   itemRepo,
		weaponRepo: weaponRepo,
		armorRepo:  armorRepo,
		jwt:        jwt,
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

func (is *ItemService) GetItemList(ctx context.Context, list_request *domain.ListDomain) (*domain.ListResponseDomain[domain.ItemDomain], error) {

	item_response, err := is.itemRepo.GetList(ctx, list_request)
	if err != nil {
		return nil, err
	}

	return item_response, nil
}
