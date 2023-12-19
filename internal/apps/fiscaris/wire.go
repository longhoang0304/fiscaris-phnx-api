package fiscaris

import (
	"github.com/google/wire"
	"github.com/longhoang0304/fiscaris-phnx-api/internal/controllers"
	"github.com/longhoang0304/fiscaris-phnx-api/internal/routes"
	"github.com/longhoang0304/fiscaris-phnx-api/libs/named_route"
)

func NewRoutes() []named_route.NamedRoute {
	pingPongController := controllers.NewPingPongController()
	pingPongRoute := routes.NewPingPongRoutes(pingPongController)

	return []named_route.NamedRoute{
		pingPongRoute,
	}
}

var InitRouter = wire.NewSet(
	NewRoutes,
	NewRouter,
)

var FiscarisAppGraphSet = wire.NewSet(
	InitRouter,
	NewHttpServer,
	NewFiscarisApp,
)
