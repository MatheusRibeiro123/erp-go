package routes

import (
	"net/http"
	"erp-go/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.Home)
}