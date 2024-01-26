package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// User represents the user for this application
// swagger:model
type User struct {
    // required: true
	gorm.Model
	// the firstname for this user
    // required: true
    // min length: 3
	Firstname  string    `json:"firstname"`
	// the lastname for this user
    // required: true
    // min length: 3
	Lastname   string    `json:"lastname"`
	// the email for this user
    // required: true
    // min length: 6
	Email      string    `json:"email"`
	// the password for this user
	Password   string    `json:"password"`
	// the role for this user
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
