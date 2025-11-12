package app

import (
	"log"
	"net/http"
	"wow-ruby/internal/config"
	"wow-ruby/internal/pkg"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	server *Server
}

func New(cfg *config.Config) (*App, error) {
	db, err := pkg.NewPostgres()
	if err != nil {
		return nil, err
	}
	log.Println("Database is connected.")

	appInstance := NewAppInstance(db, cfg.JWTSecret)

	router := chi.NewRouter()

	router.Handle("/swagger/*", httpSwagger.WrapHandler)
	RegisterAllRoutes(router, appInstance)

	srv := &Server{
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
	}
	log.Println("Server is created.")

	return &App{server: srv}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
