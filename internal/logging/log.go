package logging

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	debug = flag.Bool("debug", false, "sets log level to debug")
)

func Log() zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// 出力形式の設定
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	output.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	// ログレベルの設定
	// 通常はinfo以上を出力。-debugの場合ログレベルをdebug以上に変更する
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return log
}
