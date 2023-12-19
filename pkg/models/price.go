package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type Price struct {
	gorm.Model
	TaskId    string `json:"taskId"`
	DurationId   string `json:"durationId"`
	Amount   float64 `json:"amount"`

	Task   *Task `json:"task" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Duration   *Duration `json:"duration" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

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

func GetPriceByTaskAndDuration(taskId int64, duration string) (*Price, *gorm.DB) {
	var price Price
	db := db.Where("task_id = ? AND task_duration = ?", taskId, duration).First(&price)
	return &price, db
}


func DeletePrice(Id int64) *Price {
	var price *Price
	db.Where("ID=?", Id).Delete(&price)
	return price
}
