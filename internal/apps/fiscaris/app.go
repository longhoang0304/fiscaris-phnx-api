package fiscaris

import (
	"context"
	"github.com/longhoang0304/fiscaris-phnx-api/libs/shutdown"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type FiscarisApp struct {
	httpServer *HTTPServer
}

func (self *FiscarisApp) Start(ctx context.Context) error {
	eg, childCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return shutdown.BlockListen(childCtx, func() error {
			log.Info().Msgf("Api's listening on %s", self.httpServer.Addr)
			return self.httpServer.ListenAndServe()
		})
	})

	return eg.Wait()
}

func (self *FiscarisApp) Stop(ctx context.Context) error {
	return self.httpServer.Shutdown(ctx)
}

func NewFiscarisApp(httpServer *HTTPServer) *FiscarisApp {
	return &FiscarisApp{
		httpServer,
	}
}
