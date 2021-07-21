package connect_database

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenServer() (db *gorm.DB, err error) {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
}
