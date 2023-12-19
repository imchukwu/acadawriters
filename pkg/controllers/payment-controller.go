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


func CreatePayment(w http.ResponseWriter, r *http.Request) {
	
	CreatePayment := &models.Payment{}
	utils.ParseBody(r, CreatePayment)
	payment := CreatePayment.CreatePayment()
	res, _ := json.Marshal(payment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	payments := models.GetPayments()
	res, _ := json.Marshal(payments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentId := vars["paymentId"]
	Id, err := strconv.ParseInt(paymentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	payment, _ := models.GetPayment(Id)
	res, _ := json.Marshal(payment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentId := vars["paymentId"]
	Id, err := strconv.ParseInt(paymentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	payment := models.DeletePayment(Id)
	res, _ := json.Marshal(payment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var updatePayment = &models.Payment{}
	utils.ParseBody(r, updatePayment)

	vars := mux.Vars(r)
	paymentId := vars["paymentId"]
	Id, err := strconv.ParseInt(paymentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	payment, db := models.GetPayment(Id)
	if payment.Amount != "" {
		payment.Amount = updatePayment.Amount
	}
	if payment.Status != "" {
		payment.Status= updatePayment.Status
	}
	
	db.Save(&payment)

	res, _ := json.Marshal(payment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
