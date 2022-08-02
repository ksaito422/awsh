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

	// 操作対象のリソースとアクションを選択して、メインの処理は各パッケージで実行
	select_action := route.Main()
	controller.Main(cfg, select_action)
}
