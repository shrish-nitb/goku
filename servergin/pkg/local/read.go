package local

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (List TodoList) Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := json.Marshal(List)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.Data(http.StatusOK, "application/json", response)
	}
}
