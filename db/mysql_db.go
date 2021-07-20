package db

import (
	"github.com/jinzhu/gorm"
	//	表示只用执行init函数
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

const URL = "root:123456@tcp(192.168.56.100:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func InitMySQL() (db *gorm.DB, err error) {
	dsn := URL
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return DB, err
	}
	//	设置单数表
	DB.SingularTable(true)

	return DB, err
}

func Close() {
	_ = DB.Close()
}
