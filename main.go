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

	repo := &repositories.ClientRepository{DB: database.DB}
	service := &services.ClientService{Repository: repo}
	handler := &handlers.ClientHandler{Service: service}

	router := gin.Default()

	routes.LoadClientRoutes(router, handler)

	fmt.Println("Server is running on port 8080...")

	err = router.Run(":8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
