package apperrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	switch {
	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})

	case errors.Is(err, ErrDuplicateKey):
		c.JSON(http.StatusConflict, gin.H{"error": "duplicate key"})

	case errors.Is(err, ErrForeignKey):
		c.JSON(http.StatusConflict, gin.H{"error": "foreign key violation"})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})

	}
}
