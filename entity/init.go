package entity

import (
	"log"

	"github.com/jinzhu/gorm"

	// Mysql dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB exported to be used in models for running database queries
var DB *gorm.DB

// Mysql is used to mysql connection
func Mysql(user, password, host, dbname, port string) {
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Could not connect to database with error: " + err.Error())
	}
	db.DB()
	db.AutoMigrate(&User{})
	DB = db
}

// Redis is used to redis connection
func Redis() {
	log.Println("Connect Redis")
}
