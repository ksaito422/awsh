package main

import (
	"awsh/internal/controller"
	"awsh/internal/route"
	"awsh/internal/welcome"
	"awsh/pkg/config"
)

func main() {
	welcome.Main()
	cfg := config.Cfg()

	// Select resources and actions to be manipulated, and controller the main process.
	select_action := route.Main()
	controller.Main(cfg, select_action)
}
