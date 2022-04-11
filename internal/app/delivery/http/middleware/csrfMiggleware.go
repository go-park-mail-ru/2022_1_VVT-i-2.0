package middleware

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// /*
// 	AllowedHeaders: []string{"access-control-allow-origin", "content-type",
// 		"x-csrf-token", "access-control-expose-headers"},
// */

// func GetCsrfConfig() middleware.CSRFConfig {
// 	// TODO: это все заголовки?
// 	return middleware.CSRFConfig{
// 		CookieMaxAge: 3 * 86400,
// 		Skipper: func(context echo.Context) bool {
// 			if context.Request().RequestURI == confiu login" ||
// 				context.Request().RequestURI == "/api/v1/sendcode" ||
// 				context.Request().RequestURI == "/signup" {
// 				return true
// 			}
// 			return false
// 		},
// 	}
// }
