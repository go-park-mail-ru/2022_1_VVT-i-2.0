package serv

import (
	"net/http"

	handlers "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/handlers"
	middleware "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/middleware"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	router *mux.Router
	logger *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logger,
	}
	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	noAuthRequiredRouter := s.router.PathPrefix("/api/v1").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants", handlers.RestaurantsHandler).Methods(http.MethodGet, http.MethodOptions)
	noAuthRequiredRouter.HandleFunc("/register", handlers.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)
	noAuthRequiredRouter.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	noAuthRequiredRouter.Use(middleware.AuthOptMiddleware)

	authRequiredRouter := s.router.PathPrefix("/api/v1/auth").Subrouter()
	authRequiredRouter.HandleFunc("/logout", handlers.LogoutHandler).Methods(http.MethodPost, http.MethodOptions)
	authRequiredRouter.Use(middleware.AuthRequiredMiddleware)

	middleWareLogger := middleware.Logger{Logger: s.logger}
	s.router.Use(middleware.CorsMiddleware)
	s.router.Use(middleWareLogger.AccessLogMiddleware)
	s.router.Use(middleWareLogger.PanicMiddleware)
}
