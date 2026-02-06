package gormfx

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg Config, db *sql.DB, log logger.Interface) (*gorm.DB, error) {
	if cfg.Dialect != "mysql" && cfg.Dialect != "mariadb" {
		return nil, fmt.Errorf("%w: unsupported dialect: %s", ErrInvalidConfig, cfg.Dialect)
	}

	orm, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: log})

	if err != nil {
		return nil, fmt.Errorf("open connection: %w", err)
	}

	return orm, nil
}
