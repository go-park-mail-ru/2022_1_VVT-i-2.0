package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
	AllowedHeaders: []string{"access-control-allow-origin", "content-type",
		"x-csrf-token", "access-control-expose-headers"},
*/

func GetCorsConfig(allowOrigins []string) middleware.CORSConfig {
	// TODO: это все заголовки?
	return middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.POST, echo.OPTIONS, echo.PUT},
		ExposeHeaders:    []string{echo.HeaderXCSRFToken},
		// ExposeHeaders:    []string{echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		MaxAge: 86400,
	}
}
