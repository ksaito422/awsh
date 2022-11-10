package main

import (
	"os"

	"awsh/internal/endpoints"
	"awsh/internal/logging"
	"awsh/pkg/api/config"
)

func main() {
	endpoints.Welcome()
	cfg := config.Cfg()

	// Select resources and actions to be manipulated, and controller the main process.
	action := endpoints.Operation()
	if err := endpoints.Controller(cfg, action); err != nil {
		log := logging.Log()
		log.Debug().Stack().Err(err)
	}

	os.Exit(0)
}
