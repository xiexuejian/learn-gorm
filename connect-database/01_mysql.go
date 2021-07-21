package connect_database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//_ "example.com/my_mysql_driver" 自定义的
)

// 	参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
func OpenMysql() (db *gorm.DB, err error) {

	//	注意：想要正确的处理 time.Time ，您需要带上 parseTime 参数
	//	要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
	//	更多参数，查看官方文档：https://github.com/go-sql-driver/mysql#parameters
	//	关于字符集：https://mathiasbynens.be/notes/mysql-utf8mb4
	//	更多配置：https://github.com/go-gorm/mysql
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("打开失败")
	}

	return db, err
}

//	高级配置
func OpenMysql2() (db *gorm.DB, err error) {

	const URL = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       URL,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
}

//	自定义驱动
//	GORM 允许通过 DriverName 选项自定义 MySQL 驱动
func OpenMysql3() (db *gorm.DB, err error) {
	// Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	return gorm.Open(mysql.New(mysql.Config{
		//	自定义的驱动my_mysql_driver
		DriverName: "my_mysql_driver",
		DSN:        "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})
}

//	现有的数据库连接
//	GORM 允许通过一个现有的数据库连接来初始化 *gorm.DB
func OpenMysql4() (db *gorm.DB, err error) {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)

	return gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

}
