package routes

import (
	"erp-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LoadClientRoutes(router *gin.Engine, handler *handlers.ClientHandler) {
	router.GET("/clients", handler.GetAll)
	router.GET("/clients/:id", handler.GetByID)
	router.POST("/clients", handler.Create)
	router.PUT("/clients/:id", handler.Update)
	router.DELETE("/clients/:id", handler.Delete)
}
