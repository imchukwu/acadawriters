package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

// swagger:model Contact
type Contact struct {
	gorm.Model
	// Name of the user
   	// in: string
	Name    string `json:"name"`
	// Email of the user
   	// in: string
	Email   string `json:"email"`
	// Phone number of the user
   	// in: string
	Phone   string `json:"phone"`
	// Contact Message by user
   	// in: string
	Message string `json:"message"`
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
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
