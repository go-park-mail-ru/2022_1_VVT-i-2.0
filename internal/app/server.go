package serv

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
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
	// TODO: set prefix "api/v1" anywhere

	noAuthRequiredRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants/{city}/{page:[0-9]+}", hello).Methods(http.MethodGet)
	noAuthRequiredRouter.HandleFunc("/login", authHandler).Methods(http.MethodPost)
	noAuthRequiredRouter.Use(s.authOptMiddleware)

	authRequiredRouter := s.router.PathPrefix("/auth").Subrouter()
	authRequiredRouter.HandleFunc("/h", hello)
	// authRequiredRouter.Use(s.RequiredAuthMiddleware)
	authRequiredRouter.Use(s.authOptMiddleware)

	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
}

func hello(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(keyUserId)
	if userId != nil {
		fmt.Println("\nhello, %s", userId)
		return
	}
	fmt.Println("\nhello, incognito")
}

// s.router.HandleFunc("/restorants/{city}/{page_num}", getRestaurants).Methods(http.MethodGet)
// getRestarants нужно вернуть список ресторанов в данном городе, установить куку города
// инфа обработчика: колво ресторанов на странице
