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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /users GetUsers
    //
    // Returns all users 
	//
	// Insert Documentation
    // ---
    // produces:
    // - application/json
    // responses:
    //   '200':
    //     description: user response
    //     schema:
    //       type: array
    //       items:
    //         "$ref": "#/definitions/User"
	users := models.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail, _ := models.GetUser(Id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /users createUsers
    // ---
    // produces:
    // - application/json
    // parameters:
    //   - name: Body
    //     in: query
    //     schema:
    //       "$ref": "#/definitions/User"
    // responses:
    //   '200':
    //     description: user response
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)

	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetail, db := models.GetUser(Id)
	if updateUser.Firstname != "" {
		userDetail.Firstname = updateUser.Firstname
	}
	if updateUser.Lastname != "" {
		userDetail.Lastname = updateUser.Lastname
	}
	if updateUser.Email != "" {
		userDetail.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		userDetail.Password = updateUser.Password
	}
	
	db.Save(&userDetail)

	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
