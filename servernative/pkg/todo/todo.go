package todo

import (
	"log"
	"net/http"
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
	todoHandler.Use(func(ctx *httpserverapp.Context, w http.ResponseWriter, r *http.Request) {
		log.Println("Request came to todo handler ", r.Method, r.URL.RequestURI())
		w.Write([]byte("todolist"))
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
