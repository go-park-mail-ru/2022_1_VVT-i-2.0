package configRouting

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func ConfigureCommonMiddleware(router *echo.Echo, mwChain *middleware.CommonMiddlewareChain, corsCfg *config.CorsConfig, csrfCfg *config.CsrfConfig) {
	router.HTTPErrorHandler = mwChain.ErrorHandler
	router.Use(mwChain.PanicMiddleware)
	router.Use(mwChain.RequestIdMiddleware)
	router.Use(mwChain.AccessLogMiddleware)
	router.Use(mwChain.AuthMiddleware)
	router.Use(echoMiddleware.CORSWithConfig(getCorsConfig(corsCfg)))
	router.Use(echoMiddleware.CSRFWithConfig(getCsrfConfig(csrfCfg)))
	router.Use(middleware.CsrfSetHeader)
}
