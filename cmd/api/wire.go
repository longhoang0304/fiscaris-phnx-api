//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/longhoang0304/fiscaris-phnx-api/internal/apps/fiscaris"
)

func InitFiscarisApp() (*fiscaris.FiscarisApp, error) {
	wire.Build(fiscaris.FiscarisAppGraphSet)
	return &fiscaris.FiscarisApp{}, nil
}
