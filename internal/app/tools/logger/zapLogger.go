package logger

import (
	"fmt"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(cfg conf.LogConfig) (*zap.SugaredLogger, error) {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("unknown logging level")
	}
	zapCfg := zap.Config{
		Level:            level,
		Encoding:         cfg.Encoding,
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: cfg.ErrorOutputPaths,

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     cfg.MessageKey,
			LevelKey:       cfg.LevelKey,
			TimeKey:        cfg.TimeKey,
			CallerKey:      cfg.CallerKey,
			FunctionKey:    cfg.FunctionKey,
			StacktraceKey:  cfg.StacktraceKey,
			NameKey:        cfg.NameKey,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
		},
	}
	loggerUnsugared, err := zap.Config(zapCfg).Build()
	if err != nil {
		return nil, errors.Wrap(err, "creating ZapLogger failed")
	}

	return loggerUnsugared.Sugar(), nil
}
