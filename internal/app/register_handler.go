package serv

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"
)

// TODO: добавить валидацию данных
// TODO: добавить в jwt инфу про устройстро и страну для безопасности 3 лекция 2ч50
func hasSuchUserPhone(phone string) bool {
	for dataToLogin, _ := range usersDataBase {
		if dataToLogin.Phone == phone {
			return true
		}
	}
	return false
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	type registerData struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Username string `json:"name"`
	}
	type registerResponse struct {
		Username string `json:"username,omitempty"`
		UserAddr string `json:"userAddress,omitempty"`
		Err      string `json:"error,omitempty"`
	}

	user := r.Context().Value(keyUser)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&registerResponse{Err: "already authorized"})
	}

	requestRegisterData := &registerData{}
	if err := json.NewDecoder(r.Body).Decode(requestRegisterData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&registerResponse{Err: "wrong register data"})
		fmt.Println(err)
		return
	}

	fmt.Print("\nBODY-DATA:\n")
	fmt.Println(requestRegisterData)

	if hasSuchUserPhone(requestRegisterData.Phone) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "such user already exists"})
	}

	token, err := createToken(idIncrement, requestRegisterData.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&registerResponse{Err: "failed to create user"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "failed to register user"})
		return
	}

	idIncrement++
	usersDataBase[loginData{Phone: requestRegisterData.Phone, Password: requestRegisterData.Password}] = userDataStruct{id: idIncrement, name: requestRegisterData.Username}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, +3),
	}

	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(&LoginResponse{Username: requestRegisterData.Username})
}
