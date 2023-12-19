package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Amount   string `json:"amount"`
	Status   string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Payment{})
}

func (payment *Payment) CreatePayment() *Payment {
	db.Create(&payment)
	return payment
}

func GetPayments() []Payment {
	var payments []Payment
	db.Find(&payments)
	return payments
}

func GetPayment(Id int64) (*Payment, *gorm.DB) {
	var payment Payment
	db := db.Where("ID=?", Id).Find(&payment)
	return &payment, db
}

func DeletePayment(Id int64) *Payment {
	var payment *Payment
	db.Where("ID=?", Id).Delete(&payment)
	return payment
}
