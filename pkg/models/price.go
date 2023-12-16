package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type TaskPricing struct {
	gorm.Model
	TaskId    string `json:"taskId"`
	Timeline   string `json:"timeline"`
	Price   string `json:"price"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Price{})
}

func (price *Price) CreatePrice() *Price {
	db.Create(&price)
	return price
}

func GetPrices() []Price {
	var prices []Price
	db.Find(&prices)
	return prices
}

func GetPrice(Id int64) (*Price, *gorm.DB) {
	var price Price
	db := db.Where("ID=?", Id).Find(&price)
	return &price, db
}

func DeletePrice(Id int64) *Price {
	var price *Price
	db.Where("ID=?", Id).Delete(&price)
	return price
}
