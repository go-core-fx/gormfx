package gormfx

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(dialect gorm.Dialector, log logger.Interface) (*gorm.DB, error) {
	orm, err := gorm.Open(dialect, &gorm.Config{Logger: log})

	if err != nil {
		return nil, fmt.Errorf("open connection: %w", err)
	}

	return orm, nil
}
