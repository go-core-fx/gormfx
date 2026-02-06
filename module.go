package gormfx

import (
	"github.com/go-core-fx/logger"
	"go.uber.org/fx"
	glogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

const ModuleName = "gormfx"

func Module() fx.Option {
	return fx.Module(
		ModuleName,
		logger.WithNamedLogger(ModuleName),
		fx.Provide(NewLogger),
		fx.Provide(New),
		fx.Invoke(func(l glogger.Interface) {
			if gl, ok := l.(*zapgorm2.Logger); ok {
				gl.SetAsDefault()
			}
		}),
	)
}
