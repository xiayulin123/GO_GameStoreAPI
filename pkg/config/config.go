package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connection() {
	database, err := gorm.Open("mysql", "root:1234567@tcp(127.0.0.1:8080)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
