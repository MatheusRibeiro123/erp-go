package main

import ("fmt"
        "log"
        "net/http"

        "erp-go/internal/routes"
        "erp-go/internal/database"

        "github.com/joho/godotenv"
        )

func main() {
    
    err := godotenv.Load()
   
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    database.Connect()
    
    routes.LoadRoutes()
    
    fmt.Println("Server is running on port 8080...")
    
    err = http.ListenAndServe(":8080", nil)
    
    if err != nil {
        log.Fatal("Error starting server:", err)
    }
}