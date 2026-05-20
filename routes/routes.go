package routes

import (
	"net/http"
	"erp-go/handlers"
)

func LoadRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/products", handlers.Products)
	http.HandleFunc("/clients", handlers.Clients)
}