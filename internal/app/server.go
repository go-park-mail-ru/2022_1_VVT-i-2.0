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

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	noAuthRequiredRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRequiredRouter.HandleFunc("/restaurants", restaurants)
	// noAuthRequiredRouter.Use(s.AuthMiddleware)

	authRequiredRouter := s.router.PathPrefix("/auth").Subrouter()
	authRequiredRouter.HandleFunc("/h", hello)
	// authRequiredRouter.Use(s.RequiredAuthMiddleware)

	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Println("hello")
}

// s.router.HandleFunc("/restorants/{city}/{page_num}", getRestaurants).Methods(http.MethodGet)
// getRestarants нужно вернуть список ресторанов в данном городе, установить куку города
// инфа обработчика: колво ресторанов на странице
