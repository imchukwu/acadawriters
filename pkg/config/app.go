package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/acadawriter?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "ba15d8eef097fd:bc99a3bc@tcp(us-cdbr-east-06.cleardb.net:3306)/heroku_60b5ede111b594a?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
