package middleware

import (
	auth "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
)

type CommonMiddlewareChain struct {
	AllowOrigins []string
	Logger       *log.ServLogger
	AuthManager  auth.AuthManager
}

func NewCommonMiddlewareChain(logger *log.ServLogger, authManager auth.AuthManager, allowOrigins []string) CommonMiddlewareChain {
	return CommonMiddlewareChain{Logger: logger,
		AllowOrigins: allowOrigins,
		AuthManager:  authManager,
	}
}
