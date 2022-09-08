run:
	go run cmd/*.go
build:
	go build -o awsh cmd/main.go
mod:
	go mod tidy
doc: # localhost:6060
	godoc
