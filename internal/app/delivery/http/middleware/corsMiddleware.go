package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

func getCorsConfig(allowOrigins []string) middleware.CORSConfig {
	// TODO: это все заголовки?
	return middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
		ExposeHeaders:    []string{"authorization", "x-csrf-token"},
		MaxAge:           10000,
	}
}
