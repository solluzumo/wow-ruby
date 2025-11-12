package service

import (
	"context"
	"errors"
	"wow-ruby/internal/domain"
	"wow-ruby/internal/pkg"
	"wow-ruby/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
	jwt  string
}

func NewUserService(repo repository.UserRepository, jwt string) *UserService {
	return &UserService{
		repo: repo,
		jwt:  jwt,
	}
}

func (us *UserService) GetUserById(ctx context.Context, id string) (*domain.UserDomain, error) {
	data, err := us.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		ID:    data.ID,
		Login: data.Login,
		Hash:  data.Hash,
	}, nil
}

func (us *UserService) GetUserByLogin(ctx context.Context, login string) (*domain.UserDomain, error) {

	data, err := us.repo.FindByLogin(ctx, login)

	if err != nil {
		return nil, err
	}

	return &domain.UserDomain{
		ID:       data.ID,
		Login:    data.Login,
		Password: data.Hash,
	}, nil

}

func (us *UserService) CreateUser(ctx context.Context, userData *domain.UserDomain) (string, error) {

	hash, err := pkg.HashPassword(userData.Password)

	userData.Hash = hash

	if err != nil {
		return "Не могу создать хэш", err
	}

	userExists, _ := us.GetUserByLogin(ctx, userData.Login)
	if userExists != nil {
		return "Пользователь уже существует", errors.New("пользователь уже существует")
	}

	userId, err := us.repo.Create(ctx, userData)
	if err != nil {
		return "Не удалось создать пользователя", err
	}

	return userId, nil
}

func (us *UserService) LoginUser(ctx context.Context, userData *domain.UserDomain) (map[string]string, error) {

	user, err := us.GetUserByLogin(ctx, userData.Login)
	if user == nil {
		return nil, err
	}

	isPassword := pkg.ComparePassword(user.Hash, userData.Password)
	if isPassword != nil {
		return nil, errors.New("неверный пароль")
	}

	tokens, err := pkg.GenerateToken(userData.Login)
	if (tokens["acess_token"] == "") || (err != nil) {
		return nil, err
	}

	return tokens, nil
}
