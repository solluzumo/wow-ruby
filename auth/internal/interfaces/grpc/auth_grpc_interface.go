package grpcInterface

import (
	"context"

	"github.com/solluzumo/wow-ruby/auth/internal/domain"
	"github.com/solluzumo/wow-ruby/auth/internal/service"
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
)

type AuthGrpcServer struct {
	authpb.UnimplementedAuthServiceServer
	userService *service.UserService
}

func NewAuthGrpcServer(us *service.UserService) *AuthGrpcServer {
	return &AuthGrpcServer{
		userService: us,
	}
}

func (s *AuthGrpcServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	userDomain := &domain.UserDomain{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := s.userService.LoginUser(ctx, userDomain)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		Tokens: &authpb.Tokens{
			AcessToken:   response.AcessToken,
			RefreshToken: response.RefreshToken,
		},
	}, nil
}

func (s *AuthGrpcServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {

	data := domain.UserDomain{
		ID:       "",
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := s.userService.CreateUser(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Tokens: &authpb.Tokens{
			AcessToken:   response.AcessToken,
			RefreshToken: response.RefreshToken,
		},
	}, nil
}

func (s *AuthGrpcServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return nil, nil
}

func (s *AuthGrpcServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	return nil, nil
}

func (s *AuthGrpcServer) InternalToken(ctx context.Context, req *authpb.InternalTokenRequest) (*authpb.InternalTokenResponse, error) {
	return nil, nil
}

func (s *AuthGrpcServer) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	return nil, nil
}

func (s *AuthGrpcServer) UserDetail(ctx context.Context, req *authpb.UserDetailsRequest) (*authpb.UserDetailsResponse, error) {

	user, err := s.userService.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &authpb.UserDetailsResponse{
		UserId: user.ID,
		Email:  user.Email,
		Hash:   user.Hash,
	}, nil
}
