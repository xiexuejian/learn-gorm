package hello_gorm

import (
	log "github.com/sirupsen/logrus"
	conn "learn-gorm/db"
	"learn-gorm/model"
	"testing"
)

func TestHelloGorm(test *testing.T) {

	db, err := conn.InitMySQL()

	if err != nil {
		panic("数据库连接失败，请排查原因！")
	}

	//	添加数据
	student := &model.Student{
		Id:         1,
		Name:       "王二",
		Sex:        "男",
		Birth:      1997,
		Department: "计算机系",
		Address:    "河南省郑州市",
	}
	//create := db.Create(student)
	log.Infoln("添加数据情况:", student)

	//	----------------------------------------------------

	//	查找数据
	//	根据主键Id查找
	var stu model.Student
	//	查找一条
	db.First(&stu, 1)
	log.Infoln("根据Id查找的数据=====>", stu)
	//	根据birth查找
	stuList := make([]model.Student, 0, 10)
	//	查找一条
	//db.First(&stuList,"sex = ?", "男")
	//	查找复合条件的所有数据
	db.Find(&stuList, "sex = ?", "男")
	log.Infoln("所有数据====>", stuList)

	//	----------------------------------------------
	//	更新主键为1的数据
	db.Model(&model.Student{}).Where(1).Update("Sex", "女")

	//	---------------------------------------------
	// 删除数据
	db.Delete(&model.Student{}, 1)
	log.Infoln("删除成功")

	defer conn.Close()
}
