package main
import "fmt"
import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ERP running")}

func main() {
    http.HandleFunc("/", home)
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}