package serverginapp

import (
	"servergin/pkg/local"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	var todos = local.NewTodos()
	todosGroup := r.Group("/todos")
	{
		todosGroup.GET("/", todos.Read())
		todosGroup.POST("/", todos.Create())
		todosGroup.PUT("/", todos.Update())
		todosGroup.DELETE("/:id", todos.Delete())
	}
	r.Run(":8080")
}
