package serv

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"
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
		fmt.Println(err)
		return
	}

	fmt.Print("\nBODY-DATA:\n")
	fmt.Println(requestLoginData)

	userData, found := usersDataBase[*requestLoginData]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "no such user"})
	}

	token, err := createToken(userData.id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "failed to create user"})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, +3),
	}

	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(&LoginResponse{Username: userData.name, UserAddr: userData.address})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "not authorized"})
	}

	token.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, token)
}
