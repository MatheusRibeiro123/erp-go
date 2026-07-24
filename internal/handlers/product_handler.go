package handlers

import (
	"erp-go/internal/apperrors"
	"erp-go/internal/dto"
	"erp-go/internal/models"
	"erp-go/internal/services"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *services.ProductService
}

// handler para buscar todos os produtos
func (h *ProductHandler) GetAll(c *gin.Context) {

	products, err := h.Service.GetAll()
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, products)
}

// handler para buscar um produto por ID
func (h *ProductHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := h.Service.GetByID(idInt)

	if err != nil {
		apperrors.HandleError(c, err)
		return

	}
	c.JSON(http.StatusOK, product)
}

// handler para criar um novo produto
func (h *ProductHandler) Create(c *gin.Context) {

	var input dto.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Price:         *input.Price,
		StockQuantity: *input.StockQuantity,
	}

	id, err := h.Service.Create(product)

	if err != nil {
		apperrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// handler para atualizar um produto existente
func (h *ProductHandler) Update(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input dto.UpdateProductInput
	err = c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Price:         *input.Price,
		StockQuantity: *input.StockQuantity,
	}

	err = h.Service.Update(idInt, product)
	if err != nil {
		apperrors.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// handler para deletar um produto
func (h *ProductHandler) Delete(c *gin.Context) {

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
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// handler para atualizar parcialmente um produto
func (h *ProductHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input dto.PatchProductInput
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
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
