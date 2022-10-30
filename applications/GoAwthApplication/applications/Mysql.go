package applications

import (
	"github.com/mrkresnofatih/go-awth/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var gormMySqlInstance *gorm.DB

func GetGormMySqlInstance() *gorm.DB {
	if gormMySqlInstance == nil {
		dsn := "root:fatih0001@tcp(127.0.0.1:3306)/goawthdb?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		gormMySqlInstance = db
		return gormMySqlInstance
	}
	return gormMySqlInstance
}

func RunGormMigration() {
	db := GetGormMySqlInstance()
	db.AutoMigrate(&entities.Player{})
	db.AutoMigrate(&entities.Developer{})
}
