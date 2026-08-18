package handlers

import (
	"erp-go/internal/dto"
	"erp-go/internal/models"
	"erp-go/internal/responses"
	"erp-go/internal/services"
	"net/http"
	"strconv"

	"erp-go/internal/apperrors"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	Service *services.ClientService
}

// handler para buscar todos os clientes
func (h *ClientHandler) GetAll(c *gin.Context) {

	limit := 10 // valor padrão

	limitStr := c.Query("limit")
	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)

		if err != nil || limitInt < 1 || limitInt > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}

		limit = limitInt
	}

	page := 1 // valor padrão

	pageStr := c.Query("page")
	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)

		if err != nil || pageInt < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
			return
		}

		page = pageInt
	}

	clients, total, err := h.Service.GetAll(page, limit)

	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	responses.Paginated(c, clients, page, limit, total)
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
		apperrors.HandleError(c, err)
		return
	}

	responses.Success(c, "", client)
}

// handler para criar um novo cliente
func (h *ClientHandler) Create(c *gin.Context) {

	var input dto.CreateClientInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
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
		apperrors.HandleError(c, err)
		return
	}

	responses.Created(c, "Client created successfully", gin.H{"id": id})
}

// handler para atualizar um cliente existente
func (h *ClientHandler) Update(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input dto.UpdateClientInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	client := models.Client{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
	}

	err = h.Service.Update(idInt, client)

	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	responses.Success(c, "Client updated successfully", nil)
}

// handler para deletar um cliente existente
func (h *ClientHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.Service.Delete(idInt)

	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	responses.NoContent(c)
}

// handler para atualizar parcialmente um cliente existente
func (h *ClientHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input dto.PatchClientInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = h.Service.Patch(idInt, input)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	responses.Success(c, "Client updated partially successfully", nil)
}
