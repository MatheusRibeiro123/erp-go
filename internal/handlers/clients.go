package handlers

import (
	"fmt"
	"net/http"
)

// parei aqui, vou criar o handler de clients
func Clients(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lista de clientes")
}
