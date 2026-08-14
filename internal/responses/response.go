package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func Paginated(c *gin.Context, data interface{}, page int, limit int, total int) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Data: data,
		Pagination: Pagination{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	})
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
