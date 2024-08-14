package db

import (
	"time"

	"github.com/agilesoftgrowth/gommon/logger"

	"gorm.io/gorm"
)

func NewDatabase(
	logger logger.LoggerService,
	vendor DBVendor,
	name string,
	host string,
	port string,
	user string,
	password string,
	maxOpenConns int,
	maxIdleConns int,
	maxConnLifetime int,
	runMigrations bool,
	models ...any,
) (*gorm.DB, error) {
	db, err := gorm.Open(vendor.Dialector(host, port, name, user, password), &gorm.Config{})
	if err != nil {
		logger.Error("Cannot connect to database", "error", err.Error())
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Cannot get sql db", "error", err.Error())
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Error("Cannot ping db", "error", err.Error())
		return nil, err
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(maxConnLifetime) * time.Hour)

	if runMigrations {
		if err := db.AutoMigrate(models...); err != nil {
			logger.Error("Cannot run migrations", "error", err.Error())
			return nil, err
		}
	}

	return db, nil
}
