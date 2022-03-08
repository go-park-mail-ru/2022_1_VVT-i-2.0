package serv

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
)

// TODO: добавить валидацию данных
// TODO: добавить в jwt инфу про устройстро и страну для безопасности 3 лекция 2:50
func hasSuchUserPhone(phone string) bool {
	for dataToAuth, _ := range usersDataBase {
		if dataToAuth.phone == phone {
			return true
		}
	}
	return false
}

func signinHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("\nBODY-DATA:\n")
	fmt.Println(r.FormValue("password"))
	fmt.Println(r.FormValue("address"))
	fmt.Println(usersDataBase[dataToAuth{r.FormValue("phone"), r.FormValue("password")}])
	fmt.Println()

	user := r.Context().Value(keyUser)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "already authorized"})
	}

	newUserPhone := r.FormValue("phone")
	newUserPassword := r.FormValue("password")
	newUsername := r.FormValue("username")

	if hasSuchUserPhone(r.FormValue("phone")) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "such user already exists"})
	}

	idIncrement++
	usersDataBase[dataToAuth{phone: newUserPhone, password: newUserPassword}] = userDataStruct{id: idIncrement, name: newUsername}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": idIncrement, "username": newUsername})

	strToken, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "failed to authorize user"})
		return
	}

	// TODO: протухание добавить
	cookie := &http.Cookie{
		Name:     "token",
		Value:    strToken,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(&LoginResponse{Username: newUsername})
}

func signoutHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "not authorized"})
	}

	token.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, token)
}
