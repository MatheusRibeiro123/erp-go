package main

import (
	"fmt"
	"log"

	"erp-go/internal/database"
	"erp-go/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := gin.Default()

	routes.LoadRoutes(router)

	fmt.Println("Server is running on port 8080...")

	err = router.Run(":8080")

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
