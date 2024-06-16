package todo

import (
	"fmt"
	"log"
	"net/http"
	"servernative/pkg/httpserver/httpserverapp"
	"servernative/pkg/httpserver/httpserverapp/middleware"
	"sync"
)

type Id string
type Task struct {
	Mutex *sync.Mutex `json:"-"`
	Value string      `json:"value"`
}

type TodoMessage struct {
	Id   Id   `json:"id"`
	Task Task `json:"task"`
}

type TodoList map[Id]Task

func NewTodos() *TodoList {
	var init = make(TodoList)
	return &init
}

type Error struct {
	Message string
	Code    int
}

func (List TodoList) CreateHandlers() *httpserverapp.Handle {
	todoHandler := httpserverapp.New()
	middleware.Logger(todoHandler)

	todoHandler.Use(func(ctx *httpserverapp.Context, w http.ResponseWriter, r *http.Request) {
		log.Println("Request came to todo handler ", r.Method, r.URL.RequestURI())
		fmt.Println(ctx.Get("BODY"))
		todoHandler.Next()
	})

	// switch r.Method {
	// case "GET":
	// 	{
	// 		todoList, err := List.Read()
	// 	}
	// case "POST":
	// 	{
	// 		todoList, err := List.Create()
	// 	}
	// case "PATCH":
	// 	{
	// 		todoList, err := List.Update()
	// 	}
	// case "DELETE":
	// 	{
	// 		todoList, err := List.Delete()
	// 	}
	// default:

	// }
	//writing the response [TodoList]
	return todoHandler
}
