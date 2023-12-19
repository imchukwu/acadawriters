package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/models"
	"github.com/imchukwu/acadawriters/pkg/utils"
)


func CreateTask(w http.ResponseWriter, r *http.Request) {
	
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	task := CreateTask.CreateTask()
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetTasks()
	res, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	Id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task, _ := models.GetTask(Id)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	Id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(Id)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)

	vars := mux.Vars(r)
	taskId := vars["taskId"]
	Id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task, db := models.GetTask(Id)
	if task.Duration != "" {
		task.Duration = updateTask.Duration
	}
	if task.Category != "" {
		task.Category= updateTask.Category
	}
	if task.Status != "" {
		task.Status= updateTask.Status
	}
	
	db.Save(&task)

	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
