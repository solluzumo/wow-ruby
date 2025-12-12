package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/gateway/internal/domain"
	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
)

type QuestProvider interface {
}

type QuestService struct {
	quest QuestProvider
}

func NewQuestService() *QuestService {
	return &QuestService{}
}

func (qs *QuestService) GetQuestList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[domain.QuestDomain], error) {

	// quest_response, err := qs.questRepo.GetList(ctx, list_request)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (qs *QuestService) GetQuestById(ctx context.Context, id string) (*dto.QuestDetailResponse, error) {

	// 	quest, err := qs.questRepo.GetById(ctx, id)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return &dto.QuestDetailResponse{
	// 		ID:                     quest.ID,
	// 		Name:                   quest.Name,
	// 		RewardMoney:            quest.RewardMoney,
	// 		RequiredCharacterLevel: quest.RequiredCharacterLevel,
	// 		QuestLevel:             quest.QuestLevel,
	// 		Repeatable:             quest.Repeatable,
	// 		Difficulty:             quest.Difficulty,
	// 	}, nil

	return nil, nil
}
