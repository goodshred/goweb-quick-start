package dao

import (
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	db *gorm.DB
}

var mysqlDB MysqlDB

func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlDB = MysqlDB{db: db}
}

func GetMysqlDB() *gorm.DB {
	return mysqlDB.db
}
