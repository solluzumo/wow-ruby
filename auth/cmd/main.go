package main

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/solluzumo/wow-ruby/auth/internal/app"
	"github.com/solluzumo/wow-ruby/auth/internal/config"
	"github.com/solluzumo/wow-ruby/auth/internal/domain"
	grpcInterface "github.com/solluzumo/wow-ruby/auth/internal/interfaces/grpc"
	"github.com/solluzumo/wow-ruby/auth/internal/repository/postgres"
	"github.com/solluzumo/wow-ruby/auth/internal/service"
	"github.com/solluzumo/wow-ruby/pkg"
	authpb "github.com/solluzumo/wow-ruby/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	//конфиг
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("неудалось создать конфиг:%v", err)
	}
	//Создаём канал для задачи хеширования пароля
	taskChan := make(chan domain.HashTaskDomain, cfg.TaskChanCap)
	var wg sync.WaitGroup

	//создаём app для di
	app := app.NewAppInstance(&wg, &taskChan)

	//логгер
	logger := pkg.NewZapLogger()
	defer logger.Sync()

	//слушаем по tcp
	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("неудалось слушать:%v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(pkg.ZapUnaryInterceptor(logger)),
	)

	//подключаем бд
	db, err := pkg.NewPostgres()
	if err != nil {
		log.Fatalf("не подключились к бд:%v", err)
	}
	log.Println("База данных подключена.")

	//все необходимые слои для юзера
	userRepo := postgres.NewPostgresUserRepo(db)
	userService := service.NewUserService(userRepo, cfg.ArgonParams, app.UserChannels, logger)
	authServer := grpcInterface.NewAuthGrpcServer(userService)

	//Запускаем воркеров
	for i := 1; i <= cfg.HashWorkersCount; i++ {
		app.WG.Add(1)
		go service.HashWorker(i, app.HashTaskChannel, userService, app.WG)
	}
	log.Println("Создан пулл воркеров.")

	authpb.RegisterAuthServiceServer(s, authServer)

	log.Println("Auth Grpc сервер слушает на :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не получилось запустить сервер : %v", err)
	}

}
