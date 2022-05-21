package middleware

import (
	auth "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager"
	log "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
)

type CommonMiddleware struct {
	Logger      *log.ServLogger
	AuthManager auth.AuthManager
}

func NewCommonMiddleware(logger *log.ServLogger, authManager auth.AuthManager) CommonMiddleware {
	return CommonMiddleware{Logger: logger,
		AuthManager: authManager,
	}
}
