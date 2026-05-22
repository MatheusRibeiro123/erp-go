package handlers

import "fmt"
import "net/http"

func Clients(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lista de clientes")
}
