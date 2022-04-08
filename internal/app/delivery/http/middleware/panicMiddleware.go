package middleware

import (
	"fmt"
	"net/http"

	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/labstack/echo/v4"
)

// TODO: отправить ответ
func (mw *CommonMiddlewareChain) PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				requestId := GetRequestIdFromCtx(ctx)
				mw.Logger.Errorw(
					"panic recovered",
					log.ReqIdTitle, requestId,
					log.RemoteAddrTitle, ctx.Request().RemoteAddr,
					log.UrlTitle, ctx.Request().URL.Path,
					log.ErrorMsgTitle, fmt.Sprint(err),
				)
				ctx.JSON(http.StatusInternalServerError, struct {
					Error string `json:"error"`
				}{Error: "internal server error"})
			}
		}()
		return next(ctx)

	}
}
