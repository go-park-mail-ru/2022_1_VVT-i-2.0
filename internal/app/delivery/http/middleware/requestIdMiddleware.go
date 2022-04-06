package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

const RequestIdCtxKey = "reqId"

var requestId uint64 = 1

func nextRecId() uint64 {
	requestId++
	return requestId
}

func (mw *CommonMiddlewareChain) RequestIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("in reqid-mw")
		currReqId := nextRecId()
		ctx.Set(RequestIdCtxKey, currReqId)
		return next(ctx)
	}
}

func GetRequestIdFromCtx(ctx echo.Context) uint64 {
	reqId, ok := ctx.Get(RequestIdCtxKey).(uint64)
	if !ok {
		return 0
	}
	return reqId
}
