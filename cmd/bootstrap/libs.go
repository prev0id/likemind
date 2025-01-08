package bootstrap

import (
	"flag"
	"os"
	"time"

	httpin_integration "github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Deps() {
	flag.Parse()

	httpin_integration.UseGochiURLParam("path", chi.URLParam)

	msk, _ := time.LoadLocation("Europe/Moscow")

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeLocation: msk,
		TimeFormat:   time.DateTime,
	})
}
