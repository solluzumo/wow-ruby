package postgres

import (
	"wow-ruby/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresArmorRepo struct {
	*BaseRepository[models.Armor]
}

func NewPostgresArmorRepository(db *sqlx.DB) *PostgresArmorRepo {
	NewBaseRepository[models.Armor](db)
	return &PostgresArmorRepo{
		BaseRepository: NewBaseRepository[models.Armor](db),
	}
}
