package handlers

import "github.com/gin-gonic/gin"

func Products(c *gin.Context) {
	c.String(200, "Lista de produtos")
}
