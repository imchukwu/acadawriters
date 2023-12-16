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

func GetPrices(w http.ResponseWriter, r *http.Request) {
	prices := models.GetPrices()
	res, _ := json.Marshal(prices)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	priceId := vars["priceId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	price, _ := models.GetPrices(Id)
	res, _ := json.Marshal(price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePrice(w http.ResponseWriter, r *http.Request) {
	CreatePrice := &models.Price{}
	utils.ParseBody(r, CreatePrice)
	price := CreatePrice.CreatePrice()
	res, _ := json.Marshal(price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	priceId := vars["priceId"]
	Id, err := strconv.ParseInt(priceId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	price := models.DeletePrice(Id)
	res, _ := json.Marshal(price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	var updatePrice = &models.Price{}
	utils.ParseBody(r, updatePrice)

	vars := mux.Vars(r)
	priceId := vars["priceId"]
	Id, err := strconv.ParseInt(priceId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	price, db := models.GetPrice(Id)
	if price.Timeline != "" {
		price.Timeline = updateprice.Timeline
	}
	if price.Price != "" {
		price.Price = updateprice.Price
	}
	

	db.Save(&price)

	res, _ := json.Marshal(price)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}