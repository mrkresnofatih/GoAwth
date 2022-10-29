package applications

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var gormMySqlInstance *gorm.DB

func GetGormMySqlInstance() *gorm.DB {
	if gormMySqlInstance == nil {
		dsn := "root:fatih0001@tcp(127.0.0.1:3306)/GoAwthDb?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		gormMySqlInstance = db
		return gormMySqlInstance
	}
	return gormMySqlInstance
}
