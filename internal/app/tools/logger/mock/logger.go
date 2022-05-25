package mock

import (
	"github.com/stretchr/testify/mock"
)

type Logger struct {
	mock.Mock
}


func (*Logger) Debugw(msg string, keysAndValues ...interface{}) {}
func (*Logger) Errorw(msg string, keysAndValues ...interface{}) {}
func (*Logger) Fatalw(msg string, keysAndValues ...interface{}) {}
func (*Logger) Infow(msg string, keysAndValues ...interface{})  {}
func (*Logger) Panicw(msg string, keysAndValues ...interface{}) {}
func (*Logger) Warnw(msg string, keysAndValues ...interface{})  {}
func (*Logger) Sync() error                                     { return nil }

