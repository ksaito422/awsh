run:
	go run ./...
build:
	go build -o awsh ./main.go
mod:
	go mod tidy
doc: # localhost:6060
	godoc
