package mysql

import (
	Mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn_mysql = "root:669988@tcp(211.71.76.189:3306)/postman?charset=utf8mb4&parseTime=True&loc=Local"

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open(Mysql.Open(dsn_mysql), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
