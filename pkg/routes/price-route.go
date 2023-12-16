package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/onehour-backend/pkg/controllers"
)

var RegisterPriceRoutes = func(r *mux.Router) {
	r.HandleFunc("/price", controllers.CreatePrice).Methods("POST")
	r.HandleFunc("/price", controllers.GetPrices).Methods("GET")
	r.HandleFunc("/price/{priceId}}", controllers.GetPrice).Methods("GET")
	r.HandleFunc("/price/{priceId}", controllers.UpdatePrice).Methods("PUT")
	r.HandleFunc("/price/{priceId}", controllers.DeletePrice).Methods("DELETE")
}
