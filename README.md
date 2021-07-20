# learn-gorm
学习gorm的案例

# 学习环境
- IDE：idea2020.2（需要安装go插件）
- go： 1.16.5
- 使用Go module

# 说明
本仓库根据官方教程，写出的案例
<br/> 
官方文档：https://gorm.io/zh_CN/docs/index.html 

# 数据库SQL脚本
在文件夹名称为sql下

# 目录说明

项目目录结构是按照官方文档进行划分

│  go.mod 　　　　　　　依赖版本
<br/>│  go.sum   
<br/>│  main.go　　　　　main方法
<br/>│  README.md　　　　项目说明
<br/>│
<br/>├─dao
<br/>│  ├─associate 　　　　　　　　官网关联
<br/>│  ├─crud      　　　　　　　CRUD接口
<br/>│  ├─hello-gorm  　　　　　hello world案例
<br/>│  ├─senior     　　　　　 高级主题
<br/>│  └─tutorial   　　　　　 教程
<br/>├─db
<br/>│      mysql_db.go　　　　　　数据库配置
<br/>│
<br/>├─model
<br/>│      score-model.go  　　　　分数表score映射模型
<br/>│      student-model.go 　　　　学生表student映射模型
<br/>│
<br/>└─sql
        test.sql 　　　　数据库表结构和数据