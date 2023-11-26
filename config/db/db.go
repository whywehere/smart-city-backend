package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() error {
	gormDB, err := gorm.Open(mysql.Open("root:cq194068971@(127.0.0.1:3306)/smart_city?charset=utf8mb4&parseTime=True&loc=Local"))
	DB = gormDB
	return err
}
