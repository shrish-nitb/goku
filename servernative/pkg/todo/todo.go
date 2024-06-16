package todo

import (
	"fmt"
	"log"
	"servernative/pkg/httpserver/httpserverapp"
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
	todoHandler.Use(httpserverapp.HandlerFunc(func(h *httpserverapp.Handle) {
		log.Println("Request came to todo handler middleware")
		h.Next()
	}))

	getTodoHandler := httpserverapp.New()
	getTodoHandler.Use(httpserverapp.HandlerFunc(func(h *httpserverapp.Handle) {
		log.Println("Request came to post handler function ", h.Request.Method, h.Request.URL.RequestURI())
		fmt.Println(h.Context.Get("BODY"))
	}))

	todoHandler.AddRouter(httpserverapp.Pattern{Target: "POST /todo", Handle: getTodoHandler})

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
