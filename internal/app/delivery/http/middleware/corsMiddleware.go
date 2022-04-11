package middleware

import (
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
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		ExposeHeaders:    []string{"authorization", "x-csrf-token"},
		MaxAge:           86400,
	}
}
