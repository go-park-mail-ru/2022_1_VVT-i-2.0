package configRouting

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCorsConfig(cfg *config.CorsConfig) middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     cfg.AllowOrigins,
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.POST, echo.OPTIONS, echo.PUT},
		ExposeHeaders:    []string{echo.HeaderXCSRFToken},
		MaxAge:           cfg.MaxAge,
	}
}
