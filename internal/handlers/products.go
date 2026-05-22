package handlers

import "fmt"
import "net/http"

func Products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lista de produtos")
}
