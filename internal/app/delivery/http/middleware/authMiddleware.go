package middleware

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	_ "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/labstack/echo/v4"
)

const UserCtxKey = "user"
const TokenKeyCookie = "token"

type UserCtx struct {
	Id models.UserId
}

func (mw *CommonMiddlewareChain) AuthOptMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("in auth-opt-mw")
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return next(ctx)
		}
		fmt.Println("token-cookie: ", tokenCookie.Value)

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return next(ctx)
		}
		fmt.Println(payload)

		if _, err = govalidator.ValidateStruct(payload); err != nil {
			return next(ctx)
		}

		ctx.Set(UserCtxKey, UserCtx{Id: payload.Id})
		return next(ctx)
	}
}

func (mw *CommonMiddlewareChain) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("in auth-mw")
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
		}

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, httpErrDescr.BAD_AUTH_TOKEN)
		}

		if _, err = govalidator.ValidateStruct(payload); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, httpErrDescr.BAD_AUTH_TOKEN)
		}

		ctx.Set(UserCtxKey, UserCtx{Id: payload.Id})
		return next(ctx)
	}
}

func GetUserFromCtx(ctx echo.Context) *UserCtx {
	fmt.Println("============")
	fmt.Println(ctx)
	// fmt.Println(ctx.Get(UserCtxKey))
	user, ok := ctx.Get(UserCtxKey).(UserCtx)
	if !ok {
		return nil
	}
	return &user
}
