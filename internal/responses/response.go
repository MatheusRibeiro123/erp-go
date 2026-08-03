package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Message: message,
		Data:    data,
	})
}
func Success(c *gin.Context, message string, data interface{}) {
	Response(c, http.StatusOK, message, data)
}
func Created(c *gin.Context, message string, data interface{}) {
	Response(c, http.StatusCreated, message, data)
}
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
