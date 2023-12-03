.PHONY: run, fmt, build
run:
	go run  cmd/gatekeper/main.go
fmt:
	go fmt cmd/gatekeper/main.go
	go fmt internal/service/product/service.go
build:
	go build -o bot.exe cmd/gatekeper/main.go