package connect_database

import (
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPgsql() (db *gorm.DB, err error) {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})

}

//	默认情况下，它会启用 prepared statement 缓存，我们通过下面的方式去禁用
func OpenPgsql2() (db *gorm.DB, err error) {

	return gorm.Open(postgres.New(postgres.Config{

		DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		// 禁用
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

}

//	自定义驱动
//	GORM 允许通过 DriverName 选项自定义 PostgreSQL 驱动
func OpenPgsql3() (db *gorm.DB, err error) {
	return gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        "host=project:region:instance user=postgres dbname=postgres password=password sslmode=disable",
	}))
}
