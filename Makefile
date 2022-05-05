.PHONY: build
build:
	go build -o app ./cmd/main.go
	go build -o auth ./cmd/auth/main.go