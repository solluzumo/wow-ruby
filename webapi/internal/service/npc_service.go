package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
)

type NpcService struct {
	npcRepo repository.NpcRepository
}

func NewNpcService(npcRepo repository.NpcRepository) *NpcService {
	return &NpcService{
		npcRepo: npcRepo,
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
