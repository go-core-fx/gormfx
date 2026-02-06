package gormfx

import (
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func NewLogger(l *zap.Logger) logger.Interface {
	log := zapgorm2.New(l)
	log.LogLevel = logger.Info
	if l.Level() == zap.DebugLevel {
		log.LogLevel = logger.Info + 1
	}

	return &log
}
