module github.com/go-park-mail-ru/2022_1_VVT-i-2.0

go 1.17

require github.com/BurntSushi/toml v1.0.0

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	go.uber.org/zap v1.21.0
// serv v1.0.0
// serv/models v1.0.0
)

// replace serv v1.0.0 => ./internal/serv

// replace serv/models v1.0.0 => ./internal/app/models
