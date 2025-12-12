package app

import (
	grpcInterface "github.com/solluzumo/wow-ruby/ruby-api/internal/interfaces/grpc"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/repository/postgres"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/service"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

type AppInstance struct {
	Repos      *AppRepos
	Services   *AppServices
	APIServers *AppAPIServers
}

type AppAPIServers struct {
	ItemAPIServer  grpcInterface.ItemGrpcServer
	QuestAPIServer grpcInterface.QuestGrpcServer
	NpcAPIServer   grpcInterface.NpcGrpcServer
}

type AppServices struct {
	ItemService  *service.ItemService
	QuestService *service.QuestService
	NpcService   *service.NpcService
}

type AppRepos struct {
	ItemRepository  *postgres.PostgresItemRepo
	WeaponRepostory *postgres.PostgresWeaponRepo
	ArmorRepository *postgres.PostgresArmorRepo
	QuestRepository *postgres.PostgresQuestRepo
	NpcRepository   *postgres.PostgresNpcRepo
}

func NewAppInstance(db *sqlx.DB, logger *zap.Logger) *AppInstance {
	repos := &AppRepos{
		ItemRepository:  postgres.NewPostgresItemRepostiory(db),
		WeaponRepostory: postgres.NewPostgresWeaponRepository(db),
		ArmorRepository: postgres.NewPostgresArmorRepository(db),
		QuestRepository: postgres.NewPostgresQuestRepo(db),
		NpcRepository:   postgres.NewPostgresNpcRepo(db),
	}
	services := &AppServices{
		ItemService:  service.NewItemService(repos.ItemRepository, repos.WeaponRepostory, repos.ArmorRepository),
		QuestService: service.NewQuestService(repos.QuestRepository),
		NpcService:   service.NewNpcService(repos.NpcRepository),
	}

	return &AppInstance{
		Repos:    repos,
		Services: services,
	}
}
