package main

import (
	"flag"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/postgresqlx"

	orderHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/delivery/grpc"
	orderRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/repository"
	orderUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/usecase"

	orderProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
)

func main() {
	configPath := flag.String("config", "../../config/order.toml", "path to config file")
	flag.Parse()

	config := conf.NewOrderMicroserviceConfig()
	err := conf.ReadConfigFile(*configPath, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading config"))
	}

	pgxManager, err := postgresqlx.NewPostgresqlX(&config.DatabaseConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	orderRepo := orderRepo.NewOrderRepo(pgxManager)

	orderUcase := orderUcase.NewOrderUcase(orderRepo)

	orderHandler := orderHandler.NewOrderHandler(orderUcase)

	lis, err := net.Listen("tcp", config.OrderServConfig.BindAddr)
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()

	orderProto.RegisterOrderServiceServer(server, orderHandler)

	server.Serve(lis)
}
