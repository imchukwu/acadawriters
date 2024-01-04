package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type Duration struct {
	gorm.Model
	Name    string `json:"name"`
	NumberOfDays   string `json:"number_of_days"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Duration{})
}

func (duration *Duration) CreateDuration() *Duration {
	db.Create(&duration)
	return duration
}

func GetDurations() []Duration {
	var durations []Duration
	db.Find(&durations)
	return durations
}

func GetDuration(Id int64) (*Duration, *gorm.DB) {
	var duration Duration
	db := db.Where("ID=?", Id).Find(&duration)
	return &duration, db
}

func DeleteDuration(Id int64) *Duration {
	var duration *Duration
	db.Where("ID=?", Id).Delete(&duration)
	return duration
}
