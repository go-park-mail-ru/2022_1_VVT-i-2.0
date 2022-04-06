package middleware

import (
	auth "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CommonMiddlewareChain struct {
	AllowOrigins []string
	Logger       log.Logger
	AuthManager  auth.AuthManager
}

func NewCommonMiddlewareChain(logger log.Logger, authManager auth.AuthManager, allowOrigins []string) CommonMiddlewareChain {
	return CommonMiddlewareChain{Logger: logger,
		AllowOrigins: allowOrigins,
		AuthManager:  authManager,
	}
}

func (mwChain *CommonMiddlewareChain) ConfigureCommonMiddleware(router *echo.Echo) {
	router.HTTPErrorHandler = mwChain.ErrorHandler
	router.Use(mwChain.PanicMiddleware)
	router.Use(mwChain.RequestIdMiddleware)
	router.Use(mwChain.AccessLogMiddleware)
	router.Use(middleware.CORSWithConfig(getCorsConfig(mwChain.AllowOrigins)))
	router.Use(mwChain.AuthOptMiddleware)
}
