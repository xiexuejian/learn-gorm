package pool

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type Database struct {
	MaxConn int
	MaxOpen int
}

//	设置全局的引用型指针变量
var DatabaseConfig = new(Database)

//	数据库连接池
func GetConfig() *Database {
	viper.SetConfigFile("./settings.yml")
	content, err := ioutil.ReadFile("./settings.yml")
	if err != nil {
		fmt.Println("ioutil获取配置文件失败！")
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println("viper获取配置文件失败！")
	}
	cfgDatabase := viper.Sub("datasource")
	DatabaseConfig = InitDatabase(cfgDatabase)
	return DatabaseConfig
}

func InitDatabase(cfg *viper.Viper) *Database {
	db := &Database{
		MaxConn: cfg.GetInt("maxConn"),
		MaxOpen: cfg.GetInt("maxOpen"),
	}
	return db
}

func TestPool2(test *testing.T) {
	//获得一个*grom.DB对象
	DB, err := gorm.Open("mysql", "username:password@/database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Gorm 异常：", err)
	}
	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb := DB.DB()
	defer sqlDb.Close()
	database := GetConfig()
	fmt.Println("maxConn: ", database.MaxConn)
	fmt.Println("maxOpen: ", database.MaxOpen)

	//设置最大连接数
	sqlDb.SetMaxIdleConns(database.MaxConn)

	//设置最大的空闲连接数
	sqlDb.SetMaxOpenConns(database.MaxOpen)

	//获得当前的SQL配置情况
	data, _ := json.Marshal(sqlDb.Stats())
	fmt.Println(string(data))
}
