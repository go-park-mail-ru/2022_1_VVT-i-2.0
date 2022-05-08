package middleware

import (
	auth "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager"
	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
)

type CommonMiddlewareChain struct {
	Logger      *log.ServLogger
	AuthManager auth.AuthManager
}

func NewCommonMiddlewareChain(logger *log.ServLogger, authManager auth.AuthManager) CommonMiddlewareChain {
	return CommonMiddlewareChain{Logger: logger,
		AuthManager: authManager,
	}
}
