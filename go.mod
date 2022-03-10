module github.com/go-park-mail-ru/2022_1_VVT-i-2.0

go 1.17

require github.com/BurntSushi/toml v1.0.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dongri/phonenumber v0.0.0-20220127125919-1e58a2b4cf97
	github.com/gorilla/mux v1.8.0
	github.com/labstack/echo/v4 v4.7.0
	github.com/stretchr/testify v1.7.0
	github.com/zmb3/gogetdoc v0.0.0-20190228002656-b37376c5da6a // indirect
	go.uber.org/zap v1.21.0
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/tools v0.1.9 // indirect
// serv v1.0.0
// serv/models v1.0.0
)

// replace serv v1.0.0 => ./internal/serv

// replace serv/models v1.0.0 => ./internal/app/models
