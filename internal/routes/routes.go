package routes

import (
	"erp-go/internal/handlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/products", handlers.Products)
	http.HandleFunc("/clients", handlers.Clients)
}
