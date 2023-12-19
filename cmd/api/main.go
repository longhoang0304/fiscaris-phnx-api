package api

import (
	"context"
	"github.com/longhoang0304/fiscaris-phnx-api/libs/shutdown"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(shutdown.NewCtx())
	app, err := InitFiscarisApp()

	if err != nil {
		log.Error().Msgf("%v", err)
		return
	}

	eg.Go(func() error {
		return app.Start(ctx)
	})

	defer func() {
		app.Stop(context.Background())
		log.Info().Msg("The duckapon is shutting down")
	}()

	if err := eg.Wait(); err != nil {
		log.Error().Msgf("%v", err)
	}
}
