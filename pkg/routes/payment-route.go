package routes

import (
	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/controllers"
)

var RegisterPaymentRoutes = func(r *mux.Router) {
	r.HandleFunc("/payment", controllers.CreatePayment).Methods("POST")
	r.HandleFunc("/payment", controllers.GetPayments).Methods("GET")
	r.HandleFunc("/payment/{paymentId}}", controllers.GetPayment).Methods("GET")
	r.HandleFunc("/payment/{paymentId}", controllers.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payment/{paymentId}", controllers.DeletePayment).Methods("DELETE")
}
