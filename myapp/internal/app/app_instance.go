package app

import (
	"wow-ruby/internal/handlers"
	"wow-ruby/internal/repository/postgres"
	"wow-ruby/internal/service"

	"github.com/jmoiron/sqlx"
)

type AppInstance struct {
	Repos    *AppRepos
	Services *AppServices
	Handlers *AppHandlers
}

type AppHandlers struct {
	UserHandler  *handlers.UserHandler
	ItemHandler  *handlers.ItemHandler
	QuestHandler *handlers.QuestHandler
	NpcHandler   *handlers.NpcHandler
}

type AppServices struct {
	UserService  *service.UserService
	ItemService  *service.ItemService
	QuestService *service.QuestService
	NpcService   *service.NpcService
}

type AppRepos struct {
	UserRepository  *postgres.PostgresUserRepo
	ItemRepository  *postgres.PostgresItemRepo
	WeaponRepostory *postgres.PostgresWeaponRepo
	ArmorRepository *postgres.PostgresArmorRepo
	QuestRepository *postgres.PostgresQuestRepo
	NpcRepository   *postgres.PostgresNpcRepo
}

func NewAppInstance(db *sqlx.DB, jwt string) *AppInstance {
	repos := &AppRepos{
		UserRepository:  postgres.NewPostgresUserRepo(db),
		ItemRepository:  postgres.NewPostgresItemRepostiory(db),
		WeaponRepostory: postgres.NewPostgresWeaponRepository(db),
		ArmorRepository: postgres.NewPostgresArmorRepository(db),
		QuestRepository: postgres.NewPostgresQuestRepo(db),
		NpcRepository:   postgres.NewPostgresNpcRepo(db),
	}
	services := &AppServices{
		UserService:  service.NewUserService(repos.UserRepository, jwt),
		ItemService:  service.NewItemService(repos.ItemRepository, repos.WeaponRepostory, repos.ArmorRepository, jwt),
		QuestService: service.NewQuestService(repos.QuestRepository, jwt),
		NpcService:   service.NewNpcService(repos.NpcRepository, jwt),
	}
	handlers := &AppHandlers{
		UserHandler:  handlers.NewUserHandler(services.UserService),
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
