package handlers

import (
	"erp-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	Service *services.ClientService
}

func (h *ClientHandler) GetAll(c *gin.Context) {

	clients, err := h.Service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}
