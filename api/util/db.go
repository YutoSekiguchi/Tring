package util

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db gorm.DB

func InitDb() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("MYSQL_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	Db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB")

	return Db
}