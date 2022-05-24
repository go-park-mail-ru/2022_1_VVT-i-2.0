package middleware

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (mw *CommonMiddleware) ErrorHandler(err error, ctx echo.Context) {
	requestId := GetRequestIdFromCtx(ctx)

	if mw.Logger != nil && requestId > 0 {
		mw.Logger.Error(requestId, err.Error())
	}

	switch err := errors.Cause(err).(type) {
	case *echo.HTTPError:
		_ = ctx.JSON(err.Code, struct {
			Error string `json:"error"`
		}{Error: err.Message.(string)})
	default:
		_ = ctx.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "internal server error"})
	}
}
