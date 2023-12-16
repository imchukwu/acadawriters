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

func GetContacts(w http.ResponseWriter, r *http.Request) {
	
	contacts := models.GetContacts()
	res, _ := json.Marshal(contacts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contactDetail, _ := models.GetContact(Id)
	res, _ := json.Marshal(contactDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	
	CreateContact := &models.Contact{}
	utils.ParseBody(r, CreateContact)
	contact := CreateContact.CreateContact()
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contact := models.DeleteContact(Id)
	res, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	
	var updateContact = &models.Contact{}
	utils.ParseBody(r, updateContact)

	vars := mux.Vars(r)
	contactId := vars["contactId"]
	Id, err := strconv.ParseInt(contactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	contactDetail, db := models.GetContact(Id)
	if updateContact.Name != "" {
		contactDetail.Name = updateContact.Name
	}
	if updateContact.Email != "" {
		contactDetail.Email = updateContact.Email
	}
	if updateContact.Phone != "" {
		contactDetail.Phone = updateContact.Phone
	}
	if updateContact.Message != "" {
		contactDetail.Message = updateContact.Message
	}

	db.Save(&contactDetail)

	res, _ := json.Marshal(contactDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
