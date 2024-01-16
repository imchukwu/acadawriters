package config

// package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"github.com/xo/dburl"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect() {
	mysqlURL, ok := os.LookupEnv("MYSQL_URL")
	if !ok {
		log.Fatal("MYSQL_URL not found")
	}

	u, err := dburl.Parse(mysqlURL + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	d, err := gorm.Open(mysql.Open(u.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Adjust logging level as needed
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully")

	db = d
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil.")
	}
	return db
}

