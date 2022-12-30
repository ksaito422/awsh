run:
	go run ./...
tags:
	git describe --tags --abbrev=0
build:
	go build  -ldflags "-X main.version=$$(git describe --tags --abbrev=0)" -o awsh ./main.go
test:
	go test ./... -cover
mod:
	go mod tidy
doc: # localhost:6060
	godoc
