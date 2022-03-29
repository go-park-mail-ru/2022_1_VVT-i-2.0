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
	if mw.Logger != nil && requestId > 0 {
		mw.Logger.Errorw("error happent",
			log.ReqIdTitle, requestId,
			log.ErrorMsgTitle, err.Error(),
		)
	}

	switch err := errors.Cause(err); err.(type) {
	case *echo.HTTPError:
		ctx.JSON(err.(*echo.HTTPError).Code, struct {
			Body string
		}{Body: err.(*echo.HTTPError).Error()})
	default:
		ctx.JSON(500, struct {
			Body string
		}{Body: err.Error()})
	}

	// было в примере, но вроде тут не нужно
	// err = ctx.HTML(http.StatusInternalServerError, "internal")
	// if err != nil {
	// 	mw.Logger.Errorf("failed to write 500 internal after error: %s", err)
	// }
}

func ErrorHandler(err error, ctx echo.Context) {
	fmt.Println("+++++++++++++++error handler++++++++++++")
	ctx.Logger().Errorf("error happened while processing request: %s", err)

	switch err := errors.Cause(err); err.(type) {
	case *echo.HTTPError:
		ctx.JSON(err.(*echo.HTTPError).Code, struct {
			Body string
		}{Body: "internal"})
	default:
		ctx.JSON(500, struct {
			Body string
		}{Body: "internal"})
	}

	err = ctx.HTML(http.StatusInternalServerError, "internal")
	if err != nil {
		ctx.Logger().Errorf("failed to write 500 internal after error: %s", err)
	}
}
