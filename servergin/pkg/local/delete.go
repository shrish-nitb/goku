package local

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (List TodoList) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		task, exist := List[Id(c.Param("id"))]
		if !exist {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
			return
		}

		task.Mutex.Lock()
		defer task.Mutex.Unlock()

		delete(List, Id(c.Param("id")))

		response, err := json.Marshal(List)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.Data(http.StatusOK, "application/json", response)
	}
}
