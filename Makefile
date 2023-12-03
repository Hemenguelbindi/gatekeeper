.PHONY: run, fmt
run:
	go run  cmd/gatekeper/main.go
fmt:
	go fmt cmd/gatekeper/main.go
	go fmt internal/service/product/service.go
