package logger

import (
	"time"
)

type Logger interface {
	// Error(msg string, reqId uint64, errorMsg string)
	Debugw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Sync() error
}

type ServLogger struct {
	Logger Logger
}

func NewServLogger(logger Logger) *ServLogger {
	return &ServLogger{
		Logger: logger,
	}
}

const (
	AccessMsg          = "access"
	ReqIdTitle         = "request_id"
	MethodTitle        = "method"
	RemoteAddrTitle    = "remote_addr"
	UrlTitle           = "url"
	ProcesingTimeTitle = "processing_time"
	ErrorMsgTitle      = "error_msg"
)

func (l ServLogger) Access(requestId uint64, method, remoteAddr, url string, procesingTime time.Duration) {
	l.Logger.Infow(
		AccessMsg,
		ReqIdTitle, requestId,
		MethodTitle, method,
		RemoteAddrTitle, remoteAddr,
		UrlTitle, url,
		ProcesingTimeTitle, procesingTime,
	)
}

func (l ServLogger) Error(reqId uint64, errorMsg string) {
	l.Logger.Infow(
		"error",
		ReqIdTitle, reqId,
		ErrorMsgTitle, errorMsg,
	)
}
