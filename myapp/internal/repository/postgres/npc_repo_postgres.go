package postgres

import (
	"wow-ruby/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresNpcRepo struct {
	*BaseRepository[models.Npc]
}

func NewPostgresNpcRepo(db *sqlx.DB) *PostgresNpcRepo {
	NewBaseRepository[models.Npc](db)
	return &PostgresNpcRepo{
		BaseRepository: NewBaseRepository[models.Npc](db),
	}
}
