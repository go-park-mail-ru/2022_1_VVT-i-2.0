package main

import (
	"flag"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher/memcacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification/flashcall"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/postgresqlx"

	authHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/delivery/grpc"
	authRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/repository"
	authUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/usecase"

	authProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
)

func main() {
	configPath := flag.String("config", "../../config/auth.toml", "path to config file")
	flag.Parse()

	config := conf.NewAuthMicroserviceConfig()
	err := conf.ReadConfigFile(*configPath, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading config"))
	}

	pgxManager, err := postgresqlx.NewPostgresqlX(&config.DatabaseConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	// jwtManager := jwt.NewJwtManager(config.AuthentificatorConfig)

	// if jwtManager == nil {
	// 	log.Fatal(errors.Wrap(err, "error creating jwt-manager object"))
	// }

	memcacher, err := memcacher.NewMemcacher(&config.CacherConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating memcacher"))
	}

	flashcaller := flashcall.NewFlashcaller(&config.NotificatorConfig)

	authRepo := authRepo.NewAuthRepo(pgxManager)

	authUcase := authUcase.NewAuthUcase(flashcaller, memcacher, authRepo)

	authHandler := authHandler.NewAuthHandler(authUcase)

	lis, err := net.Listen("tcp", config.AuthServConfig.BindAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()

	authProto.RegisterAuthServiceServer(server, authHandler)

	server.Serve(lis)
}
