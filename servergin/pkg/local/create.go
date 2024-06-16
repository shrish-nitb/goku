package local

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func (List TodoList) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todoMessage TodoMessage

		if err := c.ShouldBindJSON(&todoMessage); err != nil || strings.TrimSpace(todoMessage.Task.Value) == "" || strings.TrimSpace(string(todoMessage.Id)) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if _, exist := List[todoMessage.Id]; exist {
			c.JSON(http.StatusConflict, gin.H{"error": "Resource already exists"})
			return
		}

		var mu sync.Mutex
		List[todoMessage.Id] = Task{Mutex: &mu, Value: todoMessage.Task.Value}

		response, err := json.Marshal(List)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.Data(http.StatusOK, "application/json", response)
	}
}
