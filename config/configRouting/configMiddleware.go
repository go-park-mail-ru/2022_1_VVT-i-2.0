package configRouting

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func CsrfSetHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		k, ok := context.Get("csrf").(string)
		if ok {
			context.Response().Header().Add("X-CSRF-Token", k)
		}
		return next(context)
	}
}

func ConfigureCommonMiddleware(router *echo.Echo, mwChain *middleware.CommonMiddlewareChain) {
	router.HTTPErrorHandler = mwChain.ErrorHandler
	router.Use(mwChain.PanicMiddleware)
	router.Use(mwChain.RequestIdMiddleware)
	router.Use(mwChain.AccessLogMiddleware)
	router.Use(mwChain.AuthMiddleware)
	router.Use(echoMiddleware.CORSWithConfig(middleware.GetCorsConfig(mwChain.AllowOrigins)))
	router.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		Skipper: func(context echo.Context) bool {
			if context.Request().RequestURI == "/api/v1/login" ||
				context.Request().RequestURI == "/api/v1/register" ||
				context.Request().RequestURI == "/api/v1/sendcode" {
				return true
			}
			return false
		},
		CookiePath:     "/",
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure:   true,
	}),
		CsrfSetHeader,
		echoMiddleware.Secure(),
	)
}
