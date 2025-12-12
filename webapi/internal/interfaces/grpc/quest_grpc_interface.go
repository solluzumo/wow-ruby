package grpcInterface

import (
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/service"
)

type QuestGrpcServer struct {
	authpb.UnimplementedAuthServiceServer
	questService *service.QuestService
}

func NewQuestGrpcServer(qs *service.QuestService) *QuestGrpcServer {
	return &QuestGrpcServer{
		questService: qs,
	}
}
