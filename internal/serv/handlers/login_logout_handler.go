package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	middleware "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/middleware"
	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
	token "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/token"
)

var usersDataBase = map[models.LoginRequest]models.UserDataStruct{
	{"o@o.o", "1"}: {1, "Наташа", "Москва, Петровка 38"},
	{"t@t.t", "2"}: {2, "Сережа", "Москва, Ленинградский проспект, 39"},
}

var idIncrement models.UserId = models.UserId(len(usersDataBase))

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(middleware.KeyUserId)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "already authorized"})
	}

	requestLoginData := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(requestLoginData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "wrong register data"})
		return
	}

	userData, found := usersDataBase[*requestLoginData]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "no such user"})
	}

	tokenCookie, err := token.CreateTokenCookie(userData.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "failed to create user"})
		return
	}

	host, _, _ := net.SplitHostPort(r.Host)

	tokenCookie.Domain = host
	tokenCookie.Path = "/"

	http.SetCookie(w, &tokenCookie)

	json.NewEncoder(w).Encode(&models.LoginResponse{Username: userData.Name, UserAddr: userData.Address})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	host, _, _ := net.SplitHostPort(r.Host)
	token := &http.Cookie{
		Name:    "token",
		Domain:  host,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, -3),
	}
	http.SetCookie(w, token)
}
