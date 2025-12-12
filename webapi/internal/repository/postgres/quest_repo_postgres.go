package postgres

import (
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresQuestRepo struct {
	*BaseRepository[models.Quest]
}

func NewPostgresQuestRepo(db *sqlx.DB) *PostgresQuestRepo {
	NewBaseRepository[models.Quest](db)
	return &PostgresQuestRepo{
		BaseRepository: NewBaseRepository[models.Quest](db),
	}
}
