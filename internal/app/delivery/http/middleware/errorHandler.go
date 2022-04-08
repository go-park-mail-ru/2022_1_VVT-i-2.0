package middleware

import (
	"fmt"
	"net/http"

	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

// from handlers return this:
// return echo.NewHTTPError(http.StatusUnauthorized, errors.AUTH_REQUED_MSG)
func (mw *CommonMiddlewareChain) ErrorHandler(err error, ctx echo.Context) {
	fmt.Println("======in error handler=====")

	requestId := GetRequestIdFromCtx(ctx)

	// TODO: чекнуть норм ли ошибка логируется
	// TODO: логировать только в случае отрициательного кода причины ошибки
	if mw.Logger != nil && requestId > 0 {
		mw.Logger.Errorw("error happent",
			log.ReqIdTitle, requestId,
			log.ErrorMsgTitle, err.Error(),
		)
	}

	switch err := errors.Cause(err).(type) {
	case *echo.HTTPError:
		ctx.JSON(err.Code, struct {
			Error string `json:"error"`
		}{Error: err.Message.(string)})
	default:
		// TODO: залогировать неизвестную ошибку
		ctx.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "internal server error"})
	}
}
