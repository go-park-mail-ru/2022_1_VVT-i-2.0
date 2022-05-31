package middleware

import (
	"github.com/asaskevich/govalidator"
	_ "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/labstack/echo/v4"
)

const UserCtxKey = "user"
const TokenKeyCookie = "token"

type UserCtx struct {
	Id int64
}

func (mw *CommonMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)

		if err != nil {
			return next(ctx)
		}

		payload, err := mw.AuthManager.ParseToken(tokenCookie.Value)
		if err != nil {
			return next(ctx)
		}

		if _, err = govalidator.ValidateStruct(payload); err != nil {
			return next(ctx)
		}

		ctx.Set(UserCtxKey, UserCtx{Id: payload.Id})
		return next(ctx)
	}
}

func GetUserFromCtx(ctx echo.Context) *UserCtx {
	user, ok := ctx.Get(UserCtxKey).(UserCtx)
	if !ok {
		return nil
	}
	return &user
}
