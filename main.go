package main

import (
	"fmt"
	"log"

	"erp-go/internal/database"
	"erp-go/internal/handlers"
	"erp-go/internal/repositories"
	"erp-go/internal/routes"
	"erp-go/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	clientRepo := &repositories.ClientRepository{DB: database.DB}
	clientService := &services.ClientService{Repository: clientRepo}
	clientHandler := &handlers.ClientHandler{Service: clientService}

	prodRepo := &repositories.ProductRepository{DB: database.DB}
	prodService := &services.ProductService{Repository: prodRepo}
	prodHandler := &handlers.ProductHandler{Service: prodService}

	router := gin.Default()

	routes.LoadClientRoutes(router, clientHandler)
	routes.LoadProductRoutes(router, prodHandler)

	fmt.Println("Server is running on port 8080...")

	err = router.Run(":8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
