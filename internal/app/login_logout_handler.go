package serv

// package serv

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type loginData struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Username string `json:"username,omitempty"`
	UserAddr string `json:"userAddress,omitempty"`
	Err      string `json:"error,omitempty"`
}

type userDataStruct struct {
	id      uint64
	name    string
	address string
}

var usersDataBase = map[loginData]userDataStruct{
	{"o@o.o", "1"}: {1, "Наташа", "Москва, Петровка 38"},
	{"t@t.t", "2"}: {2, "Сережа", "Москва, Ленинградский проспект, 39"},
}

var idIncrement uint64 = uint64(len(usersDataBase))

func loginHandler(w http.ResponseWriter, r *http.Request) {

	type loginResponse struct {
		Username string `json:"username,omitempty"`
		UserAddr string `json:"userAddress,omitempty"`
		Err      string `json:"error,omitempty"`
	}

	user := r.Context().Value(keyUserId)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&loginResponse{Err: "already authorized"})
	}

	requestLoginData := &loginData{}
	if err := json.NewDecoder(r.Body).Decode(requestLoginData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&loginResponse{Err: "wrong register data"})
		return
	}

	userData, found := usersDataBase[*requestLoginData]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "no such user"})
	}

	tokenCookie, err := createTokenCookie(userData.id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "failed to create user"})
		return
	}

	host, port, _ := net.SplitHostPort(r.Host)
	fmt.Println(host)
	fmt.Println(port)

	tokenCookie.Domain = host
	tokenCookie.Path = "/"

	http.SetCookie(w, &tokenCookie)

	fmt.Println(userData)
	fmt.Println(userData.name)
	fmt.Println(userData.address)
	json.NewEncoder(w).Encode(&LoginResponse{Username: userData.name, UserAddr: userData.address})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	host, _, _ := net.SplitHostPort(r.Host)
	token := &http.Cookie{
		Name:   "token",
		Domain: host,
		Path:   "/",
		// Path:    "/",
		// Path:    "http://tavide.xyz:3000",
		Expires: time.Now().AddDate(0, 0, -3),
	}
	http.SetCookie(w, token)
}
