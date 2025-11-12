package postgres

import (
	"wow-ruby/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresWeaponRepo struct {
	*BaseRepository[models.Weapon]
}

func NewPostgresWeaponRepository(db *sqlx.DB) *PostgresWeaponRepo {
	NewBaseRepository[models.Weapon](db)
	return &PostgresWeaponRepo{
		BaseRepository: NewBaseRepository[models.Weapon](db),
	}
}
