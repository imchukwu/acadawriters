package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/controllers"
)

var RegisterDurationRoutes = func(r *mux.Router) {
	r.HandleFunc("/duration", controllers.CreateDuration).Methods("POST")
	r.HandleFunc("/duration", controllers.GetDurations).Methods("GET")
	r.HandleFunc("/duration/{durationId}", controllers.GetDuration).Methods("GET")
	r.HandleFunc("/duration/{durationId}", controllers.UpdateDuration).Methods("PUT")
	r.HandleFunc("/duration/{durationId}", controllers.DeleteDuration).Methods("DELETE")
}
