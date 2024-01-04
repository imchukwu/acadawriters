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


func CreateDuration(w http.ResponseWriter, r *http.Request) {
	
	CreateDuration := &models.Duration{}
	utils.ParseBody(r, CreateDuration)
	duration := CreateDuration.CreateDuration()
	res, _ := json.Marshal(duration)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDurations(w http.ResponseWriter, r *http.Request) {
	durations := models.GetDurations()
	res, _ := json.Marshal(durations)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDuration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	durationId := vars["durationId"]
	Id, err := strconv.ParseInt(durationId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	duration, _ := models.GetDuration(Id)
	res, _ := json.Marshal(duration)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDuration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	durationId := vars["durationId"]
	Id, err := strconv.ParseInt(durationId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	duration := models.DeleteDuration(Id)
	res, _ := json.Marshal(duration)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDuration(w http.ResponseWriter, r *http.Request) {
	var updateDuration = &models.Duration{}
	utils.ParseBody(r, updateDuration)

	vars := mux.Vars(r)
	durationId := vars["durationId"]
	Id, err := strconv.ParseInt(durationId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	duration, db := models.GetDuration(Id)
	if duration.Name != "" {
		duration.Name = updateDuration.Name
	}
	if duration.NumberOfDays != "" {
		duration.NumberOfDays = updateDuration.NumberOfDays
	}
	

	db.Save(&duration)

	res, _ := json.Marshal(duration)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
