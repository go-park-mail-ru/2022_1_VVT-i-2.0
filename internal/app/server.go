package serv

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
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
	s.router.HandleFunc("/hello", hello).Methods(http.MethodGet)
	// s.router.HandleFunc("/", hello).Methods("POST")
	// s.router.Use(s.accessLogMiddleware)

	// s.router.Use(s.accessLogMiddleware)
	// s.router.Use(s.logRequest)
	// s.router.Use(s.panic)
	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	authRequired := s.router.PathPrefix("/auth").Subrouter()
	authRequired.HandleFunc("/hello", hello)
	// TODO: юзать Handle для структур, реализующих интерфейсы 45:00 3 лекция

	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
	// // s.router.HandleFunc("/", hello).Methods("POST")
	// // TODO:роутинг с мидлвар через гориллу
}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	// fmt.Println("server: hello handler started")
	// defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
