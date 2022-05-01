package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config/configRouting"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	jwt "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager/jwtManager"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher/memcacher"
	servLog "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/zaplogger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification/flashcall"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/postgresqlx"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"

	authProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"

	suggsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/delivery/http"
	suggsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/repository"
	suggsUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/usecase"
	commentHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/delivery/http"
	commentRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/repository"
	commentUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments/usecase"
	dishesHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/delivery/http"
	dishesRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/repository"
	dishesUsecase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/usecase"
	orderHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/delivery/http"
	orderRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/repository"
	orderUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/usecase"
	restaurantsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/delivery/http"
	restaurantsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/repository"
	restaurantsUsecase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/usecase"
	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/repository"
	userUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/usecase"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	configPath := flag.String("config", "../config/serv.toml", "path to config file")
	flag.Parse()

	config := conf.NewConfig()
	err := conf.ReadConfigFile(*configPath, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading config"))
	}

	pgxManager, err := postgresqlx.NewPostgresqlX(&config.DatabaseCongig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	logger, err := zaplogger.NewZapLogger(&config.LoggerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating logger object"))
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Error occurred in logger sync")
		}
	}()

	servLogger := servLog.NewServLogger(logger)

	jwtManager := jwt.NewJwtManager(config.AuthentificatorConfig)

	if jwtManager == nil {
		log.Fatal(errors.Wrap(err, "error creating jwt-manager object"))
	}

	memcacher, err := memcacher.NewMemcacher(&config.CacherConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating memcacher"))
	}

	flashcaller := flashcall.NewFlashcaller(&config.NotificatorConfig)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	authGrpcConn, err := grpc.Dial(
		config.AuthMicroserverAddr,
		// grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to grpc-auth-microserver"))
	}
	defer authGrpcConn.Close()

	authorizerCli := authProto.NewAuthServiceClient(authGrpcConn)

	userRepo := userRepo.NewUserRepo(pgxManager)
	suggsRepo := suggsRepo.NewAddrRepo(pgxManager)
	orderRepo := orderRepo.NewOrderRepo(pgxManager)

	userUcase := userUcase.NewUsecase(flashcaller, memcacher, userRepo, staticManager, authorizerCli)
	suggsUcase := suggsUcase.NewAddrUsecase(suggsRepo)
	orderUcase := orderUcase.NewUsecase(orderRepo)

	userHandler := userHandler.NewUserHandler(userUcase, jwtManager, staticManager)
	suggsHandler := suggsHandler.NewSuggsHandler(suggsUcase)
	orderHandler := orderHandler.NewOrderHandler(orderUcase)

	restaurantsRepo := restaurantsRepo.NewRestaurantsRepo(pgxManager)
	restaurantsUsecase := restaurantsUsecase.NewRestaurantsUsecase(restaurantsRepo)
	restaurantsHandler := restaurantsHandler.NewRestaurantsHandler(restaurantsUsecase, staticManager)

	dishesRepo := dishesRepo.NewDishesRepo(pgxManager)
	dishesUsecase := dishesUsecase.NewDishesUsecase(dishesRepo)
	dishesHandler := dishesHandler.NewDishesHandler(dishesUsecase, staticManager)

	commentsRepo := commentRepo.NewCommentsRepo(pgxManager)
	commentsUsecase := commentUcase.NewCommentsUsecase(commentsRepo)
	commentsHandler := commentHandler.NewCommentsHandler(commentsUsecase)

	router := echo.New()

	serverRouting := configRouting.ServerHandlers{
		UserHandler:        userHandler,
		RestaurantsHandler: restaurantsHandler,
		SuggsHandler:       suggsHandler,
		OrderHandler:       orderHandler,
		DishesHandler:      dishesHandler,
		CommentsHandler: 	commentsHandler,
	}

	serverRouting.ConfigureRouting(router)

	comonMwChain := middleware.NewCommonMiddlewareChain(servLogger, jwtManager)
	configRouting.ConfigureCommonMiddleware(router, &comonMwChain, &config.CorsConfig, &config.CsrfConfig)

	httpServ := http.Server{
		Addr:         config.ServConfig.BindAddr,
		ReadTimeout:  time.Duration(config.ServConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServConfig.WriteTimeout) * time.Second,
		Handler:      router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		// if err := httpServ.ListenAndServeTLS("../localhost.crt", "../localhost.key"); err != http.ErrServerClosed {
		// if err := router.StartAutoTLS(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
