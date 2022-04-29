package configRouting

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/labstack/echo/v4"

	userHandler "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/delivery/http"
)

func getCsrfConfig(cfg *config.CsrfConfig) middleware.CSRFConfig {
	return middleware.CSRFConfig{
		CookieMaxAge: cfg.MaxAge,
		Skipper: func(context echo.Context) bool {
			if context.Request().RequestURI == v1Prefix+"send_code" ||
				context.Request().RequestURI == v1Prefix+"suggest" {
				return true
			}
			return false
		},
		SetterTokenInUnsafeMethod: func(context echo.Context) bool {
			if context.Request().RequestURI == v1Prefix+"login" ||
				context.Request().RequestURI == v1Prefix+"register" {
				return true
			}
			return false
		},
		CookieName: userHandler.CSRFCookieName,
		CookiePath: "/",
	}
}
