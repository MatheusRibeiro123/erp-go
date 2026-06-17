package routes

import (
	"erp-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LoadClientRoutes(router *gin.Engine, clientHandler *handlers.ClientHandler) {
	router.GET("/clients", clientHandler.GetAll)
	router.GET("/clients/:id", clientHandler.GetByID)
	router.POST("/clients", clientHandler.Create)
	router.PUT("/clients/:id", clientHandler.Update)
	router.DELETE("/clients/:id", clientHandler.Delete)
	router.PATCH("/clients/:id", clientHandler.Patch)
}
