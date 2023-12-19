package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/longhoang0304/fiscaris-phnx-api/internal/controllers"
	nr "github.com/longhoang0304/fiscaris-phnx-api/libs/named_route"
)

func NewPingPongRoutes(pingPongController controllers.IPingPongController) nr.NamedRoute {
	r := chi.NewRouter()
	r.Get("/", pingPongController.Pong)

	return nr.NamedRoute{
		Name:  "Pong",
		Route: r,
	}
}
