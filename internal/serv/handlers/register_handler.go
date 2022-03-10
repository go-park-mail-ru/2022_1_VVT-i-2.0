package handlers

import (
	"encoding/json"
	"net"
	"net/http"

	middleware "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/middleware"
	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
	token "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/token"
	validation "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/validation"
)

// TODO: добавить валидацию данных
// TODO: добавить в jwt инфу про устройстро и страну для безопасности 3 лекция 2ч50
func hasSuchUserPhone(phone string) bool {
	for dataToLogin := range usersDataBase {
		if dataToLogin.Phone == phone {
			return true
		}
	}
	return false
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.KeyUserId)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&models.RegisterResponse{Err: "already authorized"})
		return
	}

	dataToRegister := &models.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(dataToRegister); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.RegisterResponse{Err: "wrong register data"})
		return
	}

	if !validation.ValidatePhone(dataToRegister.Phone) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.RegisterResponse{Err: "not valid phone"})
		return
	}

	if !validation.ValidatePassword(dataToRegister.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.RegisterResponse{Err: "not valid password"})
		return
	}

	if !validation.ValidateUsername(dataToRegister.Username) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.RegisterResponse{Err: "not valid username"})
		return
	}

	if hasSuchUserPhone(dataToRegister.Phone) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "such user already exists"})
		return
	}

	idIncrement++
	usersDataBase[models.LoginRequest{Phone: dataToRegister.Phone, Password: dataToRegister.Password}] = models.UserDataStruct{Id: idIncrement, Name: dataToRegister.Username}

	tokenCookie, err := token.CreateTokenCookie(idIncrement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.LoginResponse{Err: "failed to create user"})
		return
	}

	host, _, _ := net.SplitHostPort(r.Host)
	tokenCookie.Domain = host
	tokenCookie.Path = "/"

	http.SetCookie(w, &tokenCookie)

	json.NewEncoder(w).Encode(&models.LoginResponse{Username: dataToRegister.Username})
}
