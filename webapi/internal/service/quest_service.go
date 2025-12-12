package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/dto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository"
)

type QuestService struct {
	questRepo repository.QuestRepository
}

func NewQuestService(questRepo repository.QuestRepository) *QuestService {
	return &QuestService{
		questRepo: questRepo,
	}
}

func (qs *QuestService) GetQuestList(ctx context.Context, list_request *dto.ListRequest) (*dto.ListResponse[models.Quest], error) {

	quest_response, err := qs.questRepo.GetList(ctx, list_request)
	if err != nil {
		return nil, err
	}

	return quest_response, nil
}

func (qs *QuestService) GetQuestById(ctx context.Context, id string) (*dto.QuestDetailResponse, error) {

	quest, err := qs.questRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.QuestDetailResponse{
		ID:                     quest.ID,
		Name:                   quest.Name,
		RewardMoney:            quest.RewardMoney,
		RequiredCharacterLevel: quest.RequiredCharacterLevel,
		QuestLevel:             quest.QuestLevel,
		Repeatable:             quest.Repeatable,
		Difficulty:             quest.Difficulty,
	}, nil
}
