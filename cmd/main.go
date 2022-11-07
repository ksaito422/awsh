package main

import (
	"awsh/internal/endpoints"
	"awsh/pkg/api/config"
)

func main() {
	endpoints.Welcome()
	cfg := config.Cfg()

	// Select resources and actions to be manipulated, and controller the main process.
	action := endpoints.Operation()
	endpoints.Controller(cfg, action)
}
