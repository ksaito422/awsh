package main

import (
	"flag"
	"fmt"
	"os"

	"awsh/internal/endpoints"
	"awsh/internal/logging"
	"awsh/pkg/api/config"
	ecsapi "awsh/pkg/api/ecs"
	s3api "awsh/pkg/api/s3"
	ecsser "awsh/pkg/service/ecs"
	s3ser "awsh/pkg/service/s3"
)

var (
	version string
)

func main() {
	option()

	endpoints.Welcome()
	cfg := config.Cfg()

	s3a := s3api.NewS3API()
	ecsa := ecsapi.NewECSAPI()
	s3s := s3ser.NewS3Service(s3a)
	ecss := ecsser.NewECSService(ecsa)
	r := endpoints.NewAppController(s3s, ecss)
	// Select resources and actions to be manipulated, and controller the main process.
	action := r.Operation()
	v := endpoints.Index(action)
	if err := r.Controller(cfg, v); err != nil {
		log := logging.Log()
		log.Error().Err(err).Msg("")
		// TODO: デバッグモードでスタックトレース出した方が良いかも
		log.Debug().Stack().Err(err).Msg("")
	}

	os.Exit(0)
}

func option() {
	var showVersion bool

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	// -v,-versionオプションでversionを表示する
	if showVersion {
		fmt.Println("awsh version: ", version)
		os.Exit(0)
	}
}
