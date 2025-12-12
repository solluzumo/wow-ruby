package postgres

import (
	"github.com/solluzumo/wow-ruby/auth/internal/models"

	"github.com/jmoiron/sqlx"
)

type BaseRepository[T models.Tabler] struct {
	DB *sqlx.DB
}

func NewBaseRepository[T models.Tabler](db *sqlx.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}
