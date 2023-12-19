package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role    string    `json:"role"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUser(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID=?", Id).Find(&user)
	return &user, db
}

func DeleteUser(Id int64) *User {
	var user *User
	db.Where("ID=?", Id).Delete(&user)
	return user
}
