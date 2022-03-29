package middleware

import (
	"fmt"
	"time"

	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/labstack/echo/v4"
)

const LoggerCtxKey = "logger"

func (mw *CommonMiddlewareChain) AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("in access-log-mw")
		reqId := GetRequestIdFromCtx(ctx)
		ctx.Set(LoggerCtxKey, mw.Logger)
		start := time.Now()
		next(ctx)
		log.AccessLog(&mw.Logger, reqId, ctx.Request().Method, ctx.Request().RemoteAddr, ctx.Request().URL.Path, time.Since(start))
		return nil
	}
}

func GetLoggerFromCtx(ctx echo.Context) *log.Logger {
	logger, ok := ctx.Get(LoggerCtxKey).(log.Logger)
	if !ok {
		return nil
	}
	return &logger
}
