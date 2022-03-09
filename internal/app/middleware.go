package serv

import (
	// "github.com/gorilla/mux"
	// "go.uber.org/zap"

	"fmt"
	"net/http"
	"time"
)

var allowedOrigins = []string{"", "http://localhost:3000", "http://travide.xyz:3000"}

// var allowedOrigins = []string{"", "http://127.0.0.1:8080", "http://127.0.0.1:3000", "https://bmstusssa.herokuapp.com", "https://bmstusa.ru"}

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

func (s *server) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		isAllowed := false
		for _, o := range allowedOrigins {
			if origin == o {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			s.logger.Errorf("CORS not allowed origin = ", origin)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS,HEAD")
		w.Header().Set("Access-Control-Expose-Headers",
			"Accept,Accept-Encoding,X-CSRF-Token,Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
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
				http.Error(w, `{"error":"error on server"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
