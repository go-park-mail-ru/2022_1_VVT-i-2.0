package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewZapLogger(cfg zap.Config) (*zap.SugaredLogger, error) {
	loggerUnsugared, err := zap.Config(cfg).Build()
	if err != nil {
		return nil, errors.Wrap(err, "creating ZapLogger failed")
	}

	return loggerUnsugared.Sugar(), nil
}
