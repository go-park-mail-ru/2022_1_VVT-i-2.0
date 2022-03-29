package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"
)

func getCorsConfig(allowOrigins []string) middleware.CORSConfig {
	// TODO: это все заголовки?
	fmt.Println("in get-cors-conf")
	return middleware.CORSConfig{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
		ExposeHeaders:    []string{"authorization", "x-csrf-token"},
		MaxAge:           10000,
	}
}
