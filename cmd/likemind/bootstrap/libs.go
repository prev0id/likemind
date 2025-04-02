package bootstrap

import (
	"flag"
	"os"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Libs() {
	flag.Parse()

	msk, _ := time.LoadLocation("Europe/Moscow")

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeLocation: msk,
		TimeFormat:   time.DateTime,
	})

	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
}
