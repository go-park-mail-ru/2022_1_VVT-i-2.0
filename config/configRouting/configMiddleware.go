package configRouting

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func ConfigureCommonMiddleware(router *echo.Echo, mwChain *middleware.CommonMiddlewareChain) {
	router.HTTPErrorHandler = mwChain.ErrorHandler
	router.Use(mwChain.PanicMiddleware)
	router.Use(mwChain.RequestIdMiddleware)
	router.Use(mwChain.AccessLogMiddleware)
	router.Use(echoMiddleware.CORSWithConfig(middleware.GetCorsConfig(mwChain.AllowOrigins)))

	router.Group(V1Prefix, mwChain.AuthOptMiddleware)
	router.Group(V1AuthPrefix, mwChain.AuthMiddleware)
}
