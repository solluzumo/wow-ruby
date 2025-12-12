package postgres

import (
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"

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
