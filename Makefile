.PHONY: build
build:
	go build -o app ./cmd/main.go
	go build -o auth ./cmd/auth/main.go
	go build -o order ./cmd/order/main.go