package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Contact{})
}

func (contact *Contact) CreateContact() *Contact {
	db.Create(&contact)
	return contact
}

func GetContacts() []Contact {
	var contacts []Contact
	db.Find(&contacts)
	return contacts
}

func GetContact(Id int64) (*Contact, *gorm.DB) {
	var contact Contact
	db := db.Where("ID=?", Id).Find(&contact)
	return &contact, db
}

func DeleteContact(Id int64) *Contact {
	var contact *Contact
	db.Where("ID=?", Id).Delete(&contact)
	return contact
}
