package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/gateway/internal/domain"
	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
)

type NpcProvider interface {
}

type NpcService struct {
}

func NewNpcService() *NpcService {
	return &NpcService{}
}

func (is *NpcService) GetNpcList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[domain.NpcDomain], error) {

	// npc_response, err := is.npcRepo.GetList(ctx, list_request)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (ns *NpcService) GetNpcById(ctx context.Context, id string) (*dto.NpcDetailResponse, error) {

	// npc, err := ns.npcRepo.GetById(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }

	// return &dto.NpcDetailResponse{
	// 	ID:          npc.ID,
	// 	Name:        npc.Name,
	// 	Health:      npc.Health,
	// 	Mana:        npc.Mana,
	// 	Level:       npc.Level,
	// 	Tameable:    npc.Tameable,
	// 	Faction:     npc.Faction,
	// 	Reaction:    npc.Reaction,
	// 	Location:    npc.Location,
	// 	RespawnTime: npc.RespawnTime,
	// }, nil
	return nil, nil
}
