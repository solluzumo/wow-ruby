package app

import (
	"github.com/solluzumo/wow-ruby/gateway/internal/handlers"
	"github.com/solluzumo/wow-ruby/gateway/internal/infrastructure/auth"
	"github.com/solluzumo/wow-ruby/gateway/internal/service"
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
)

type AppInstance struct {
	Repos    *AppRepos
	Services *AppServices
	Handlers *AppHandlers
	API      *AppAPIClients
}

type AppAPIClients struct {
	AuthApiClient *auth.AuthGrpcClient
}

type AppHandlers struct {
	UserHandler  *handlers.UserHandler
	ItemHandler  *handlers.ItemHandler
	QuestHandler *handlers.QuestHandler
	NpcHandler   *handlers.NpcHandler
}

type AppServices struct {
	UserService  *service.AuthService
	ItemService  *service.ItemService
	QuestService *service.QuestService
	NpcService   *service.NpcService
}

type AppRepos struct {
}

func NewAppInstance(db *sqlx.DB, grpcConn *grpc.ClientConn, logger *zap.Logger) *AppInstance {
	authServiceClient := authpb.NewAuthServiceClient(grpcConn)

	apis := &AppAPIClients{
		AuthApiClient: auth.NewAuthGrpcClient(&authServiceClient),
	}

	repos := &AppRepos{}
	services := &AppServices{
		UserService:  service.NewAuthService(apis.AuthApiClient),
		ItemService:  service.NewItemService(),
		QuestService: service.NewQuestService(),
		NpcService:   service.NewNpcService(),
	}
	handlers := &AppHandlers{
		UserHandler:  handlers.NewUserHandler(services.UserService, logger),
		ItemHandler:  handlers.NewItemHandler(services.ItemService),
		QuestHandler: handlers.NewQuestHandler(services.QuestService),
		NpcHandler:   handlers.NewNpcHandler(services.NpcService),
	}
	return &AppInstance{
		Repos:    repos,
		Services: services,
		Handlers: handlers,
	}
}
