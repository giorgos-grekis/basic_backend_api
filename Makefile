build:
	@go build -o bin/basic_backend_api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/basic_backend_api
