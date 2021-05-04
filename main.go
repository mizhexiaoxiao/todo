package main

import (
	"todo/models"
	"todo/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// func InitMysql() (DB *gorm.DB, err error) {
// 	DB, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/t_todo?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		.Println(err)
// 		return
// 	}
// 	return
// }

func main() {
	//连接数据库
	err := models.InitMysql()
	if err != nil {
		panic(err)
	}
	//程序退出时断开数据库连接
	defer models.Close()
	//程序启动自动创建模型表
	models.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func CreateATodo(todo *Todo)(err error){

// }
