package middleware

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/errorDescription"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/labstack/echo/v4"
)

const UserCtxKey = "user"
const TokenKeyCookie = "token"

type UserCtx struct {
	id models.UserId
}

func (mw *CommonMiddlewareChain) AuthOptMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return next(ctx)
		}

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return next(ctx)
		}

		// TODO: validate usetid from payload
		ctx.Set(UserCtxKey, UserCtx{id: payload.Id})
		return next(ctx)
	}
}

func (mw *CommonMiddlewareChain) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errorDescription.AUTH_REQUIRED_DESCR)
		}

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, errorDescription.BAD_AUTH_TOKEN)
		}

		// TODO: validate usetid from payload
		ctx.Set(UserCtxKey, UserCtx{id: payload.Id})
		return next(ctx)
	}
}

func GetUserFromCtx(ctx echo.Context) *UserCtx {
	user, ok := ctx.Get(UserCtxKey).(*UserCtx)
	if !ok {
		return nil
	}
	return user
}
