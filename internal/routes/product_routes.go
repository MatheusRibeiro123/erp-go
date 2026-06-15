package routes

import (
	"erp-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LoadProductRoutes(router *gin.Engine, prodHandler *handlers.ProductHandler) {
	router.GET("/products", prodHandler.GetAll)
	router.GET("/products/:id", prodHandler.GetByID)
	router.POST("/products", prodHandler.Create)
	router.PUT("/products/:id", prodHandler.Update)
	router.DELETE("/products/:id", prodHandler.Delete)
}
