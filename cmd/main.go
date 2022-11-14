package main

import (
	"os"

	"awsh/internal/endpoints"
	"awsh/internal/logging"
	"awsh/pkg/api/config"
	s3api "awsh/pkg/api/s3"
	s3ser "awsh/pkg/service/s3"
)

func main() {
	endpoints.Welcome()
	cfg := config.Cfg()

	s3a := s3api.NewS3API()
	s3s := s3ser.NewS3Service(s3a)
	r := endpoints.NewAppController(s3s)
	// Select resources and actions to be manipulated, and controller the main process.
	action := r.Operation()
	v := endpoints.Index(action)
	if err := r.Controller(cfg, v); err != nil {
		log := logging.Log()
		log.Debug().Stack().Err(err)
	}

	os.Exit(0)
}
