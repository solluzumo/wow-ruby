package postgres

import (
	"context"

	"github.com/solluzumo/wow-ruby/ruby-api/internal/domain"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresItemRepo struct {
	*BaseRepository[models.Item]
}

func NewPostgresItemRepostiory(db *sqlx.DB) *PostgresItemRepo {
	NewBaseRepository[models.Item](db)
	return &PostgresItemRepo{
		BaseRepository: NewBaseRepository[models.Item](db),
	}
}

func (br *PostgresItemRepo) GetItemById(ctx context.Context, id string) (*domain.ItemDomain, error) {
	var data models.Item

	query := `
		SELECT i.*, 
			COALESCE(q.id, 0) AS quest_id, 
			COALESCE(q.quest_name, '') AS quest_name
		FROM item AS i
		LEFT JOIN QuestItem qi ON qi.item_id = i.id
		LEFT JOIN Quest q ON q.id = qi.quest_id
		WHERE i.id = $1
	`

	err := br.DB.Get(&data, query, id)

	if err != nil {
		return nil, err
	}

	return &domain.ItemDomain{
		ID:            data.ID,
		Name:          data.Name,
		Price:         data.Price,
		RequiredLevel: data.RequiredLevel,
		MaxStack:      data.MaxStack,
		Rarity:        string(data.Rarity),
		ItemType:      string(data.ItemType),
		QuestId:       data.QuestId,
		QuestName:     data.QuestName,
	}, nil
}
