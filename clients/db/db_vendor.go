package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	VendorPsql DBVendor = iota
	VendorMysql
)

type DBVendor int

func (v DBVendor) Dialector(host, port, name, user, password string) gorm.Dialector {
	switch v {
	case VendorMysql:
		dsn := fmt.Sprintf("mysql://%s:%s@%s:%s/%s", user, password, host, port, name)
		return mysql.Open(dsn)
	default:
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", host, port, name, user, password)
		return postgres.Open(dsn)
	}
}
