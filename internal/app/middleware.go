package serv

import (
	// "github.com/gorilla/mux"
	// "go.uber.org/zap"
	"fmt"
	"net/http"
	"time"
)

func (s *server) accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		s.logger.Infow(
			"access",
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"url", r.URL.Path,
			"processing_time", time.Since(start).String(),
		)
	})
}

func (s *server) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.logger.Errorw(
					fmt.Sprint(err),
					"method", r.Method,
					"remote_addr", r.RemoteAddr,
					"url", r.URL.Path,
				)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
