package serv

import (
	"net/http"
	"time"
)

func Start(config *Config) error {
	loggerUnsugared, err := config.LogConfig.Build()
	if err != nil {
		return err
	}
	defer loggerUnsugared.Sync() // TODO: где это должно быть?
	logger := loggerUnsugared.Sugar()
	serv := NewServer(logger)

	httpServ := http.Server{
		Addr:         config.BindAddr,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		Handler:      serv,
	}

	return httpServ.ListenAndServe()
}
