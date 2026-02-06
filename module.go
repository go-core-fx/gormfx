package gormfx

import (
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
)

const ModuleName = "gormfx"

func Module() fx.Option {
	return fx.Module(
		ModuleName,
		logger.WithNamedLogger(ModuleName),
		// fx.Provide(New),
	)
}
