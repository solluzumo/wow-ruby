package grpcInterface

import (
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/service"
)

type NpcGrpcServer struct {
	authpb.UnimplementedAuthServiceServer
	npcService *service.NpcService
}

func NewNpcGrpcServer(ns *service.NpcService) *NpcGrpcServer {
	return &NpcGrpcServer{
		npcService: ns,
	}
}
