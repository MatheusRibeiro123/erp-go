package handlers

import (
	"erp-go/internal/dto"
	"erp-go/internal/models"
	"erp-go/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	Service *services.ClientService
}

// handler para buscar todos os clientes
func (h *ClientHandler) GetAll(c *gin.Context) {

	clients, err := h.Service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

// handler para buscar um cliente por ID
func (h *ClientHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client, err := h.Service.GetByID(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

// handler para criar um novo cliente
func (h *ClientHandler) Create(c *gin.Context) {

	var input dto.CreateClientInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.Client{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
	}

	id, err := h.Service.Create(client)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
