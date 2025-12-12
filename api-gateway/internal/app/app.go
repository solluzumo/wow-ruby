package app

import (
	"log"
	"net/http"

	"github.com/solluzumo/wow-ruby/gateway/internal/config"
	"github.com/solluzumo/wow-ruby/pkg"
	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	server         *Server
	AuthServiceRPC *grpc.ClientConn
}

func New(cfg *config.Config) (*App, error) {
	logger := pkg.NewLogger()
	defer logger.Sync()

	db, err := pkg.NewPostgres()
	if err != nil {
		log.Fatalf("не подключились к бд:%v", err)
	}
	log.Println("Database is connected.")

	authConn, err := pkg.NewAuthClientConn()
	if err != nil {
		log.Fatalf("не подключились к auth service:%v", err)
	}

	log.Println("Auth service is connected.")

	appInstance := NewAppInstance(db, authConn, logger)

	router := chi.NewRouter()

	router.Use(pkg.LoggingMiddleware(logger))

	router.Handle("/swagger/*", httpSwagger.WrapHandler)
	RegisterAllRoutes(router, appInstance)

	srv := &Server{
		server: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: router,
		},
	}
	log.Println("Server is created.")

	return &App{
		server:         srv,
		AuthServiceRPC: authConn,
	}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
