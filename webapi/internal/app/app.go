package app

import (
	"log"
	"net/http"

	"github.com/solluzumo/wow-ruby/pkg"
	"github.com/solluzumo/wow-ruby/ruby-api/internal/config"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	server *Server
}

func New(cfg *config.Config) (*App, error) {
	logger := pkg.NewLogger()
	defer logger.Sync()

	db, err := pkg.NewPostgres()
	if err != nil {
		log.Fatalf("не подключились к бд:%v", err)
	}
	log.Println("Database is connected.")

	appInstance := NewAppInstance(db, logger)

	router := chi.NewRouter()

	router.Use(pkg.LoggingMiddleware(logger))

	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	srv := &Server{
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
	}
	log.Println("Server is created.")

	return &App{
		server: srv,
	}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
