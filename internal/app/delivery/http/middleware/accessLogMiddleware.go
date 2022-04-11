package middleware

import (
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	"github.com/labstack/echo/v4"
)

const LoggerCtxKey = "logger"

func (mw *CommonMiddlewareChain) AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		reqId := GetRequestIdFromCtx(ctx)
		ctx.Set(LoggerCtxKey, mw.Logger)
		start := time.Now()
		result := next(ctx)
		mw.Logger.Access(reqId, ctx.Request().Method, ctx.Request().RemoteAddr, ctx.Request().URL.Path, time.Since(start))
		return result
	}
}

func GetLoggerFromCtx(ctx echo.Context) *logger.ServLogger {
	logger, ok := ctx.Get(LoggerCtxKey).(*logger.ServLogger)
	if !ok {
		return nil
	}
	return logger
}
