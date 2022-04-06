package logger

import (
	"time"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Sync() error
}

const (
	AccessMsg          = "access"
	ReqIdTitle         = "id"
	MethodTitle        = "method"
	RemoteAddrTitle    = "remote_addr"
	UrlTitle           = "url"
	ProcesingTimeTitle = "processing_time"
	ErrorMsgTitle      = "error_msg"
)

func AccessLog(l *Logger, requestId uint64, method, remoteAddr, url string, procesingTime time.Duration) {
	(*l).Infow(
		AccessMsg,
		ReqIdTitle, requestId,
		MethodTitle, method,
		RemoteAddrTitle, remoteAddr,
		UrlTitle, url,
		ProcesingTimeTitle, procesingTime,
	)
}
