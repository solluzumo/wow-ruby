package grpcInterface

import (
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/service"
)

type ItemGrpcServer struct {
	authpb.UnimplementedAuthServiceServer
	itemService *service.ItemService
}

func NewItemGrpcServer(is *service.ItemService) *ItemGrpcServer {
	return &ItemGrpcServer{
		itemService: is,
	}
}
