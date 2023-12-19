package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/controllers"
)

var RegisterUserRoutes = func(r *mux.Router) {
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")
}
