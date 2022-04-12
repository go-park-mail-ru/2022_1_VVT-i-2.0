package configRouting

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCsrfConfig(cfg *config.CsrfConfig) middleware.CSRFConfig {
	return middleware.CSRFConfig{
		CookieMaxAge: cfg.MaxAge,
		Skipper: func(context echo.Context) bool {
			// fmt.Println(context.Request().RequestURI)
			// fmt.Println(v1Prefix + "/send_code")
			// fmt.Println(context.Request().RequestURI == v1Prefix+"/send_code")
			if context.Request().RequestURI == v1Prefix+"login" ||
				context.Request().RequestURI == v1Prefix+"send_code" ||
				context.Request().RequestURI == v1Prefix+"register" {
				return true
			}
			return false
		},
		CookiePath:     "/",
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure:   true,
	}
}
