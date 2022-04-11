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

	// "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher/memcacher"
	servLog "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/zaplogger"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification/flashcall"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/postgresqlx"

	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/repository"
	userUcase "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/usecase"
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

	userRepo := userRepo.NewUserRepo(pgxManager)
	// // restaurantRepo := restaurantRepo.NewRestaurantRepository(postgresDB.GetDatabase())
	// // dishRepo := dishRepo.NewDishRepository(postgresDB.GetDatabase())

	userUcase := userUcase.NewUsecase(flashcaller, memcacher, userRepo)
	// // restaurantUcase := restaurantUsecase.NewRestaurantUsecase(restaurantRepo)
	// // dishUcase := dishUsecase.NewDishUsecase(dishRepo)

	userHandler := userHandler.NewUserHandler(userUcase, jwtManager)
	// // restaurantHandler := restaurantHandler.NewRestaurantHandler(restaurantUcase)
	// // dishHandler := dishHandler.NewDishHandler(dishUcase)

	router := echo.New()

	serverRouting := configRouting.ServerHandlers{
		UserHandler: userHandler,
		// OrderHandler: orderHandler,
		// CartHandler:  cartHandler,
		//...
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
		log.Fatal(err)
	}
}
