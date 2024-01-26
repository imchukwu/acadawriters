// Package classification Acadawriters API.
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//    
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/imchukwu/acadawriters/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterContactRoutes(r)
	routes.RegisterPriceRoutes(r)
	routes.RegisterDurationRoutes(r)
	routes.RegisterTaskRoutes(r)
	routes.RegisterPaymentRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting server %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
