package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	DB *gorm.DB
)



func InitMySQL()(err error){
	dsn := "root:mmblydblk3@tcp(127.0.0.1:3306)/todomvc?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}








