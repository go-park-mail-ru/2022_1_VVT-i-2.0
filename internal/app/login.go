package serv

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type dataToAuth struct {
	email    string
	password string
}

var usersDataBase = map[dataToAuth]int{
	dataToAuth{"o@o.o", "1"}: 1,
	dataToAuth{"t@t.t", "2"}: 2,
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := r.Context().Value(keyUserId)
	if userId != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	fmt.Println(r.FormValue("email"), r.FormValue("password"))
	fmt.Println(usersDataBase[dataToAuth{r.FormValue("email"), r.FormValue("password")}])

	userId, found := usersDataBase[dataToAuth{r.FormValue("email"), r.FormValue("password")}]
	if !found {
		w.WriteHeader(http.StatusForbidden)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId})

	strToken, err := token.SignedString(SECRET)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    strToken,
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
