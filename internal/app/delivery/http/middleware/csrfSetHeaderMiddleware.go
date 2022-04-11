package middleware

import (
	"github.com/labstack/echo/v4"
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
