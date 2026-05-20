package main
import "fmt"
import "net/http"
import "erp-go/routes"

func main() {
    routes.SetupRoutes()
    
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}