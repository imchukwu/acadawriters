package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/controllers"
)

var RegisterTaskRoutes = func(r *mux.Router) {
	r.HandleFunc("/task", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/task", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/task/{taskId}}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/task/{taskId}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{taskId}", controllers.DeleteTask).Methods("DELETE")
}
