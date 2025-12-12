package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/solluzumo/wow-ruby/gateway/internal/dto"
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
)

type AuthGrpcClient struct {
	client authpb.AuthServiceClient
}

func NewAuthGrpcClient(client *authpb.AuthServiceClient) *AuthGrpcClient {
	return &AuthGrpcClient{
		client: *client,
	}
}

func (a *AuthGrpcClient) Login(ctx context.Context, email, password string) (*dto.UserLoginResponse, error) {
	loginRequest := &authpb.LoginRequest{
		Email:    email,
		Password: password,
	}

	resp, err := a.client.Login(ctx, loginRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка от серивса авторизации: %v", err)
	}

	return &dto.UserLoginResponse{
		AcessToken:   resp.Tokens.AcessToken,
		RefreshToken: resp.Tokens.RefreshToken,
	}, nil
}

func (a *AuthGrpcClient) Register(ctx context.Context, email, password string) (*dto.UserLoginResponse, error) {
	registerRequest := &authpb.RegisterRequest{
		Email:    email,
		Password: password,
	}
	start := time.Now()
	resp, err := a.client.Register(ctx, registerRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка от серивса авторизации: %v", err)
	}

	grpcDuration := time.Since(start)

	return &dto.UserLoginResponse{
		AcessToken:   resp.Tokens.AcessToken,
		RefreshToken: resp.Tokens.RefreshToken,
		GrpcDuration: grpcDuration,
	}, nil
}

func (a *AuthGrpcClient) Logout(ctx context.Context, token string) (*dto.LogoutResponse, error) {
	logoutRequest := &authpb.LogoutRequest{
		AcessToken: token,
	}

	resp, err := a.client.Logout(ctx, logoutRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка от серивса авторизации: %v", err)
	}

	return &dto.LogoutResponse{
		Success: resp.Success,
	}, nil
}

func (a *AuthGrpcClient) UserDetail(ctx context.Context, email string) (*dto.UserDetailResponse, error) {
	userDetailRequest := &authpb.UserDetailsRequest{
		Email: email,
	}

	resp, err := a.client.UserDetail(ctx, userDetailRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка от серивса авторизации: %v", err)
	}

	return &dto.UserDetailResponse{
		ID:    resp.UserId,
		Email: resp.Email,
		Hash:  resp.Hash,
	}, nil
}
