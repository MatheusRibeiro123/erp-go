package handlers

import (
	"github.com/gin-gonic/gin"
)

func Clients(c *gin.Context) {
	c.String(200, "Lista de clientes")
}
