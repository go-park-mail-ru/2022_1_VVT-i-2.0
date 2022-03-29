package main

// читаем конфиг
// создание объекта БД
// создание объекта логера

// создание объектов-репозиториев
// создание объектов-юзкейсов
// создание объектов-http-хендлеров

// создание роутера, настрока хендлеров и миддлеваре

// err = server.ListenAndServe() с нашим роутером

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config/configRouting"

	jwt "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager/jwtManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	// userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery"
	// userRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/repository"
	// userUsecase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/usecase"
	// restaurantHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurant/delivery"
	// restaurantRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurant/repository"
	// restaurantUsecase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurant/usecase"
	// dishHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dish/delivery"
	// dishRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dish/repository"
	// dishUsecase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dish/usecase"
)

func main() {
	configPath := flag.String("config", "../config/serv.toml", "path to config file")
	flag.Parse()

	config := conf.NewConfig()
	err := conf.ReadConfigFile(*configPath, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading config"))
	}

	// var postgresConfig config.ServConfig
	// err := config.ReadConfigFile(*configPath, servConfig)
	// postgresDB, err := database.NewPostgres(servConfig.GetPostgresConfig())
	// if err != nil {
	// 	log.Fatal(errors.Wrap(error, "error creating postgres agent"))
	// }
	// defer postgresDB.Close()

	logger, err := logger.NewZapLogger(config.LoggerConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating logger object"))
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Error occurred in logger sync")
		}
	}()

	jwtManager := jwt.NewJwtManager(config.AuthConfig)

	if jwtManager == nil {
		log.Fatal(errors.Wrap(err, "error creating jwt-manager object"))
	}

	// userRepo := userRepo.NewUserRepository(postgresDB.GetDatabase())
	// // restaurantRepo := restaurantRepo.NewRestaurantRepository(postgresDB.GetDatabase())
	// // dishRepo := dishRepo.NewDishRepository(postgresDB.GetDatabase())

	// userUcase := userUsecase.NewUserUsecase(userRepo)
	// // restaurantUcase := restaurantUsecase.NewRestaurantUsecase(restaurantRepo)
	// // dishUcase := dishUsecase.NewDishUsecase(dishRepo)

	// userHandler := userHandler.NewUserHandler(userUcase)
	// // restaurantHandler := restaurantHandler.NewRestaurantHandler(restaurantUcase)
	// // dishHandler := dishHandler.NewDishHandler(dishUcase)

	router := echo.New()

	serverRouting := configRouting.ServerHandlers{
		// UserHandler: userHandler,
		// OrderHandler: orderHandler,
		// CartHandler:  cartHandler,
		//...
	}

	serverRouting.ConfigureRouting(router)

	comonMwChain := middleware.NewCommonMiddlewareChain(logger, jwtManager, config.ServConfig.AllowOrigins)
	comonMwChain.ConfigureCommonMiddleware(router)

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
