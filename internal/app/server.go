package serv

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type server struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func newServer(logger *zap.SugaredLogger) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000", "http://178.154.229.61:8080", "http://tavide.xyz:8080"}),
		handlers.AllowedHeaders([]string{
			"Accept", "Content-Type", "Content-Length",
			"Accept-Encoding", "X-CSRF-Token", "csrf-token", "Authorization"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	))

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	noAuthRequiredRouter := s.router.PathPrefix("/api/v1").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants", restaurants).Methods(http.MethodPost, http.MethodOptions)
	noAuthRequiredRouter.HandleFunc("/register", registerHandler).Methods(http.MethodPost, http.MethodOptions)
	noAuthRequiredRouter.HandleFunc("/login", loginHandler).Methods(http.MethodPost, http.MethodOptions)
	noAuthRequiredRouter.Use(s.authOptMiddleware)

	authRequiredRouter := s.router.PathPrefix("/api/v1/auth").Subrouter()
	authRequiredRouter.HandleFunc("/logout", logoutHandler).Methods(http.MethodPost, http.MethodOptions)
	authRequiredRouter.Use(s.authRequiredMiddleware)

	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
}
