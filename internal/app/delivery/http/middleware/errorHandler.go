package middleware

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

// from handlers return this:
// return echo.NewHTTPError(http.StatusUnauthorized, errors.AUTH_REQUED_MSG)
func (mw *CommonMiddlewareChain) ErrorHandler(err error, ctx echo.Context) {
	requestId := GetRequestIdFromCtx(ctx)

	// TODO: чекнуть норм ли ошибка логируется
	if mw.Logger != nil && requestId > 0 {
		mw.Logger.Error(requestId, err.Error())
	}

	switch err := errors.Cause(err).(type) {
	case *echo.HTTPError:
		ctx.JSON(err.Code, struct {
			Error string `json:"error"`
		}{Error: err.Message.(string)})
	default:
		ctx.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "internal server error"})
	}
}
