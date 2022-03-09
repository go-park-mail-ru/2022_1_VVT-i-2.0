module github.com/go-park-mail-ru/2022_1_VVT-i-2.0

go 1.13

require github.com/BurntSushi/toml v1.0.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/zmb3/gogetdoc v0.0.0-20190228002656-b37376c5da6a // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/tools v0.1.9 // indirect
	serv v1.0.0
)

replace serv v1.0.0 => ./internal/app
