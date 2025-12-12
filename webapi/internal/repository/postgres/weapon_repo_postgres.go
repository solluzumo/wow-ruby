package postgres

import (
	"github.com/solluzumo/wow-ruby/ruby-api/internal/models"

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
