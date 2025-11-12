package postgres

import (
	"wow-ruby/internal/models"

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
