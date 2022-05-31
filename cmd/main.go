package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/metrics"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config/configRouting"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	jwt "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager/jwtManager"

	servLog "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/zaplogger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/postgresqlx"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"

	authProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
	orderProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"

	suggsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/delivery/http"
	suggsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/repository"
	suggsUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/usecase"
	dishesHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/delivery/http"
	dishesRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/repository"
	dishesUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes/usecase"
	orderHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/delivery/http"
	orderUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order/usecase"
	promocodeHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode/delivery/http"
	promocodeRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode/repository"
	promocodeUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode/usecase"
	recommendationsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/delivery/http"
	recommendationsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/repository"
	recommendationsUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/usecase"
	restaurantsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/delivery/http"
	restaurantsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/repository"
	restaurantsUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants/usecase"
	reviewsHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews/delivery/http"
	reviewsRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews/repository"
	reviewsUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews/usecase"
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

	pgxManager, err := postgresqlx.NewPostgresqlX(&config.DatabaseConfig)
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

	orderGrpcConn, err := grpc.Dial(
		config.OrderMicroserverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to grpc-auth-microserver"))
	}
	defer orderGrpcConn.Close()

	orderCli := orderProto.NewOrderServiceClient(orderGrpcConn)

	userRepo := userRepo.NewUserRepo(pgxManager)
	suggsRepo := suggsRepo.NewAddrRepo(pgxManager)

	userUcase := userUcase.NewUcase(userRepo, staticManager, authorizerCli)
	suggsUcase := suggsUcase.NewAddrUcase(suggsRepo)
	orderUcase := orderUcase.NewUcase(orderCli)

	userHandler := userHandler.NewUserHandler(userUcase, jwtManager, staticManager)
	suggsHandler := suggsHandler.NewSuggsHandler(suggsUcase)
	orderHandler := orderHandler.NewOrderHandler(orderUcase, staticManager)

	restaurantsRepo := restaurantsRepo.NewRestaurantsRepo(pgxManager)
	restaurantsUcase := restaurantsUcase.NewRestaurantsUcase(restaurantsRepo)
	restaurantsHandler := restaurantsHandler.NewRestaurantsHandler(restaurantsUcase, staticManager)

	promocodeRepo := promocodeRepo.NewPromocodeRepo(pgxManager)
	promocodeUcase := promocodeUcase.NewPromocodeUcase(promocodeRepo)
	promocodeHandler := promocodeHandler.NewPromocodesHandler(promocodeUcase, staticManager)

	dishesRepo := dishesRepo.NewDishesRepo(pgxManager)
	dishesUcase := dishesUcase.NewDishesUcase(dishesRepo)
	dishesHandler := dishesHandler.NewDishesHandler(dishesUcase, staticManager)

	commentsRepo := reviewsRepo.NewReviewsRepo(pgxManager)
	commentsUcase := reviewsUcase.NewRestaurantReviewsUcase(commentsRepo)
	commentsHandler := reviewsHandler.NewRestaurantReviewsHandler(commentsUcase)

	recommendationsRepo := recommendationsRepo.NewRecommendationsRepo(pgxManager)
	recommendationsUcase := recommendationsUcase.NewRecommendationsUcase(recommendationsRepo)
	recommendationsHandler := recommendationsHandler.NewRecommendationsHandler(recommendationsUcase, staticManager)

	router := echo.New()

	m, err := metrics.CreateNewMetric("main")
	if err != nil {
		panic(err)
	}

	router.Use(m.CollectMetrics)

	serverRouting := configRouting.ServerHandlers{
		UserHandler:            userHandler,
		RestaurantsHandler:     restaurantsHandler,
		SuggsHandler:           suggsHandler,
		OrderHandler:           orderHandler,
		DishesHandler:          dishesHandler,
		ReviewsHandler:         commentsHandler,
		RecommendstionsHandler: recommendationsHandler,
		PromocodeHandler:       promocodeHandler,
	}

	comonMw := middleware.NewCommonMiddleware(servLogger, jwtManager)
	serverRouting.ConfigureRouting(router, &comonMw, &config.CorsConfig, &config.CsrfConfig)

	httpServ := http.Server{
		Addr:         config.ServConfig.BindAddr,
		ReadTimeout:  time.Duration(config.ServConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServConfig.WriteTimeout) * time.Second,
		Handler:      router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
