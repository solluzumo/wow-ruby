package postgres

import (
	"context"
	"fmt"
	"wow-ruby/internal/domain"
	"wow-ruby/internal/models"

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

func (ur *PostgresUserRepo) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE login = $1"

	err := ur.DB.Get(&user, query, login)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *PostgresUserRepo) FindByLoginAndHash(ctx context.Context, login string, hash string) (*models.User, error) {
	var user models.User
	fmt.Println(login, hash)
	query := "SELECT * FROM users WHERE login = $1 AND hash = $2"

	err := ur.DB.Get(&user, query, login, hash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *PostgresUserRepo) Create(ctx context.Context, user *domain.UserDomain) (string, error) {

	id := uuid.New().String()

	_, err := ur.DB.Exec(
		"INSERT INTO users(id,login,hash) values($1, $2, $3)",
		id,
		user.Login,
		user.Hash,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}
