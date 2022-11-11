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

	r := endpoints.NewAppController(&endpoints.Route{})
	// Select resources and actions to be manipulated, and controller the main process.
	action := r.Operation()
	if err := r.Controller(cfg, action); err != nil {
		log := logging.Log()
		log.Debug().Stack().Err(err)
	}

	os.Exit(0)
}
