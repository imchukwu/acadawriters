package config

import (
	"fmt"
	"os"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
    mysqlURL, ok := os.LookupEnv("MYSQL_URL")
    if ok == false {
        fmt.Println("MYSQL_URL not found")
        return
    }

    u, err := dburl.Parse(mysqlURL + "?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        fmt.Println(err)
        return
    }

    d, err := gorm.Open(mysql.Open(u.DSN), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("connected to database successfully")

	db = d
}

func GetDB() *gorm.DB {
	return db
}


// package config

// import (
// 	"os"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var (
// 	db *gorm.DB
// )

// func Connect() {

// 	dsn := os.Getenv("MYSQL_URL");
// 	// dsn := "root:password@tcp(127.0.0.1:3306)/acadawriters?charset=utf8mb4&parseTime=True&loc=Local"
// 	// dsn := "mysql://root:-dCB1Ad6HdHddCGH3BdEcECgggbDh4gb@roundhouse.proxy.rlwy.net:20981/acadawriters"

// 	// dsn := "ba15d8eef097fd:bc99a3bc@tcp(us-cdbr-east-06.cleardb.net:3306)/heroku_60b5ede111b594a?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db = d
// }

// func GetDB() *gorm.DB {
// 	return db
// }
