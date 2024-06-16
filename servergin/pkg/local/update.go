package local

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (List TodoList) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var todoMessage TodoMessage
		if err := c.ShouldBindJSON(&todoMessage); err != nil || strings.TrimSpace(todoMessage.Task.Value) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		task, exist := List[todoMessage.Id]
		if !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
			return
		}

		task.Mutex.Lock()
		defer task.Mutex.Unlock()

		List[todoMessage.Id] = Task{Mutex: task.Mutex, Value: todoMessage.Task.Value}

		response, err := json.Marshal(List)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.Data(http.StatusOK, "application/json", response)
	}
}
