package db

import (
	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DBParams struct {
	Logger          logger.LoggerService
	Vendor          DBVendor
	DBName          string
	Host            string
	Port            string
	User            string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime int
	RunMigrations   bool
	Models          []any
}

type DBResult struct {
	fx.Out
	DB *gorm.DB
}

func NewDB(params DBParams) (DBResult, error) {
	db, err := NewDbConnection(
		params.Logger,
		params.Vendor,
		params.DBName,
		params.Host,
		params.Port,
		params.User,
		params.Password,
		params.MaxOpenConns,
		params.MaxIdleConns,
		params.MaxConnLifetime,
		params.RunMigrations,
		params.Models...,
	)
	return DBResult{DB: db}, err
}
