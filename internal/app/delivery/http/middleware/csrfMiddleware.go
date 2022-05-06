package middleware

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

const csrfCtxKey = "csrf"

type (
	// CSRFConfig defines the config for CSRF middleware.
	CSRFConfig struct {
		Skipper                   func(ctx echo.Context) bool
		SetterTokenInUnsafeMethod func(ctx echo.Context) bool
		TokenLength               uint8
		CookieName                string
		CookieDomain              string
		CookiePath                string
		CookieMaxAge              int
	}
)

var (
	DefaultCSRFConfig = CSRFConfig{
		Skipper:                   func(c echo.Context) bool { return false },
		SetterTokenInUnsafeMethod: func(c echo.Context) bool { return false },
		TokenLength:               32,
		CookieName:                "_csrf",
		CookieMaxAge:              86400,
	}
)

func CSRF() echo.MiddlewareFunc {
	c := DefaultCSRFConfig
	return CSRFWithConfig(c)
}

func CSRFWithConfig(config CSRFConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultCSRFConfig.Skipper
	}
	if config.TokenLength == 0 {
		config.TokenLength = DefaultCSRFConfig.TokenLength
	}
	if config.CookieName == "" {
		config.CookieName = DefaultCSRFConfig.CookieName
	}
	if config.CookieMaxAge == 0 {
		config.CookieMaxAge = DefaultCSRFConfig.CookieMaxAge
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if config.Skipper(ctx) {
				return next(ctx)
			}

			token := ""
			if tokenCookie, err := ctx.Cookie(config.CookieName); err != nil || tokenCookie.Value == "" {
				token = random.String(config.TokenLength)
			} else {
				token = tokenCookie.Value
			}

			switch {
			case config.SetterTokenInUnsafeMethod(ctx), ctx.Request().Method == http.MethodGet, ctx.Request().Method == http.MethodHead, ctx.Request().Method == http.MethodOptions, ctx.Request().Method == http.MethodTrace:
			default:
				// Validate token only for requests which are not defined as 'safe' by RFC7231
				clientToken := ctx.Request().Header.Get(echo.HeaderXCSRFToken)
				if clientToken != token {
					return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.INVALID_CSRF)
				}
			}

			cookie := &http.Cookie{
				Expires:  time.Now().Add(time.Duration(config.CookieMaxAge) * time.Second),
				Secure:   true,
				HttpOnly: true,
				Name:     config.CookieName,
				Value:    token,
			}
			if config.CookiePath != "" {
				cookie.Path = config.CookiePath
			}
			if config.CookieDomain != "" {
				cookie.Domain = config.CookieDomain
			}
			ctx.SetCookie(cookie)

			ctx.Set(csrfCtxKey, token)

			ctx.Response().Header().Add(echo.HeaderVary, echo.HeaderCookie)
			return next(ctx)
		}
	}
}

func GetCSRFTokenromCtx(ctx echo.Context) string {
	token, ok := ctx.Get(csrfCtxKey).(string)
	if !ok {
		return ""
	}
	return token
}

func CsrfSetHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		k, ok := ctx.Get("csrf").(string)
		if ok && ctx.Request().Method == http.MethodGet || ctx.Request().Method == http.MethodHead || ctx.Request().Method == http.MethodOptions || ctx.Request().Method == http.MethodTrace {
			ctx.Response().Header().Add("X-CSRF-Token", k)
		}
		return next(ctx)
	}
}
