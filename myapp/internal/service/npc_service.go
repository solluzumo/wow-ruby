package service

import (
	"context"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/models"
	"wow-ruby/internal/repository"
)

type NpcService struct {
	npcRepo repository.NpcRepository
	jwt     string
}

func NewNpcService(npcRepo repository.NpcRepository, jwt string) *NpcService {
	return &NpcService{
		npcRepo: npcRepo,
		jwt:     jwt,
	}
}

func (is *NpcService) GetNpcList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[models.Npc], error) {

	npc_response, err := is.npcRepo.GetList(ctx, list_request)
	if err != nil {
		return nil, err
	}

	return npc_response, nil
}

func (ns *NpcService) GetNpcById(ctx context.Context, id string) (*dto.NpcDetailResponse, error) {

	npc, err := ns.npcRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.NpcDetailResponse{
		ID:          npc.ID,
		Name:        npc.Name,
		Health:      npc.Health,
		Mana:        npc.Mana,
		Level:       npc.Level,
		Tameable:    npc.Tameable,
		Faction:     npc.Faction,
		Reaction:    npc.Reaction,
		Location:    npc.Location,
		RespawnTime: npc.RespawnTime,
	}, nil
}
