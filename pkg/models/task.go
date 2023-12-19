package models

import (
	"github.com/imchukwu/acadawriters/pkg/config"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Category   string `json:"category"`
	Duration string `json:"duration"`
	Status   string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (task *Task) CreateTask() *Task {
	db.Create(&task)
	return task
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func GetTask(Id int64) (*Task, *gorm.DB) {
	var task Task
	db := db.Where("ID=?", Id).Find(&task)
	return &task, db
}

func DeleteTask(Id int64) *Task {
	var task *Task
	db.Where("ID=?", Id).Delete(&task)
	return task
}
