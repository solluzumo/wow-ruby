package postgres

import (
	"context"

	"github.com/solluzumo/wow-ruby/auth/internal/domain"
	"github.com/solluzumo/wow-ruby/auth/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepo struct {
	*BaseRepository[models.User]
}

func NewPostgresUserRepo(db *sqlx.DB) *PostgresUserRepo {
	NewBaseRepository[models.User](db)
	return &PostgresUserRepo{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}

func (ur *PostgresUserRepo) UpdateStatus(ctx context.Context, email string) error {
	_, err := ur.DB.Exec(
		"UPDATE users WHERE email = $1 set is_active=true",
		email,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ur *PostgresUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE email = $1"

	err := ur.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *PostgresUserRepo) Create(ctx context.Context, user *domain.UserDomain) (string, error) {

	id := uuid.New().String()

	_, err := ur.DB.Exec(
		"INSERT INTO users(id,hash,email) values($1, $2, $3)",
		id,
		user.Hash,
		user.Email,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}
