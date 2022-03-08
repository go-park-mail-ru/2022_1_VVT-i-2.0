package serv

import (
	"fmt"
	"net/http"

	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
)

type dataToAuth struct {
	phone    string
	password string
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

var idIncrement uint64 = 3
var usersDataBase = map[dataToAuth]userDataStruct{
	{"o@o.o", "1"}: {1, "Наташа", "Москва, Петровка 38"},
	{"t@t.t", "2"}: {2, "Сережа", "Москва, Ленинградский проспект, 39"},
}

func createToken(userId uint64, userAddr string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId, "userAddress": userAddr})

	return token.SignedString(SECRET)

}

func signupHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("\nBODY-DATA:\n")
	fmt.Println(r.FormValue("phone"), r.FormValue("password"))
	fmt.Println(r.FormValue("password"))
	fmt.Println(usersDataBase[dataToAuth{r.FormValue("phone"), r.FormValue("password")}])
	fmt.Println()

	user := r.Context().Value(keyUser)
	if user != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "already authorized"})
	}

	userData, found := usersDataBase[dataToAuth{r.FormValue("phone"), r.FormValue("password")}]
	if !found {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&LoginResponse{Err: "no such user"})
	}

	token, err := createToken(userData.id, userData.address)

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
	}

	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(&LoginResponse{Username: userData.name, UserAddr: userData.address})
}
