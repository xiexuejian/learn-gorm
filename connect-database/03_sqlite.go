package connect_database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite() (db *gorm.DB, err error) {
	// https://github.com/mattn/go-sqlite3
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}

//	注意： 您也可以使用 file::memory:?cache=shared 替代文件路径。 这会告诉 SQLite 在系统内存中使用一个临时数据库
//	https://www.sqlite.org/inmemorydb.html
