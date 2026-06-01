package routes

import (
	"erp-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {
	router.GET("/", handlers.Home)
	router.GET("/products", handlers.Products)
	router.GET("/clients", handlers.Clients)
}
