package serv

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	user := r.Context().Value(keyUserId)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&registerResponse{Err: "already authorized"})
	}

	dataToRegister := &registerData{}
	if err := json.NewDecoder(r.Body).Decode(dataToRegister); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&registerResponse{Err: "wrong register data"})
		return
	}

	if hasSuchUserPhone(dataToRegister.Phone) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "such user already exists"})
	}

	idIncrement++
	fmt.Println("----------------")
	fmt.Println(dataToRegister)
	usersDataBase[loginData{Phone: dataToRegister.Phone, Password: dataToRegister.Password}] = userDataStruct{id: idIncrement, name: dataToRegister.Username}
	fmt.Println(usersDataBase[loginData{Phone: dataToRegister.Phone, Password: dataToRegister.Password}])

	tokenCookie, err := createTokenCookie(idIncrement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "failed to create user"})
		return
	}

	http.SetCookie(w, &tokenCookie)

	json.NewEncoder(w).Encode(&LoginResponse{Username: dataToRegister.Username})
}
