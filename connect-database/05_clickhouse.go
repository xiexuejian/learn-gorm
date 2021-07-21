package connect_database

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type User struct {
}

//	https://github.com/go-gorm/clickhouse
func OpenClickhouse() (db *gorm.DB, err error) {
	dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
	db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接失败")
	}

	return db, err
}

func DemoClick() {

	db, _ := OpenClickhouse()

	// Auto Migrate
	_ = db.AutoMigrate(&User{})
	// Set table options
	_ = db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})

	// 插入
	db.Create(&User{})

	// 查询
	db.Find(&User{}, "id = ?", 10)

	var user1 User
	var user2 User
	var user3 User
	// 批量插入
	var users = []User{user1, user2, user3}
	db.Create(&users)
}
