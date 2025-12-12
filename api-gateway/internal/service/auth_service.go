package service

import (
	"context"

	"github.com/solluzumo/wow-ruby/gateway/internal/domain"
	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
)

type AuthProvider interface {
	Login(ctx context.Context, email, password string) (*dto.UserLoginResponse, error)
	Register(ctx context.Context, email, password string) (*dto.UserLoginResponse, error)
	UserDetail(ctx context.Context, email string) (*dto.UserDetailResponse, error)
	Logout(ctx context.Context, token string) (*dto.LogoutResponse, error)
}

type AuthService struct {
	auth AuthProvider
}

func NewAuthService(ap AuthProvider) *AuthService {
	return &AuthService{
		auth: ap,
	}
}

func (as *AuthService) Register(ctx context.Context, userData *domain.UserDomain) (*dto.UserLoginResponse, error) {

	return as.auth.Register(ctx, userData.Email, userData.Password)
}

func (as *AuthService) Login(ctx context.Context, userData *domain.UserDomain) (*dto.UserLoginResponse, error) {
	return as.auth.Login(ctx, userData.Email, userData.Password)
}

func (as *AuthService) GetByEmail(ctx context.Context, email string) (*dto.UserDetailResponse, error) {
	return as.auth.UserDetail(ctx, email)
}

func (as *AuthService) Logout(ctx context.Context, token string) (*dto.LogoutResponse, error) {
	return as.auth.Logout(ctx, token)
}
