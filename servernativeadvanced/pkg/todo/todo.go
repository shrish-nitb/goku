package todo

import (
	"log"
	"servernative/pkg/httpserver/httpserverapp"
	"sync"
	"time"
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
		log.Println("Request came to GET handler function ", h.Request.Method, h.Request.URL.RequestURI())
		time.Sleep(30 * time.Second)
		res := (h.Context.Get("BODY")).(string)
		h.Writer.Header().Set("a", "b")
		h.Writer.Write([]byte(res))
		log.Println("Response sent")
	}))

	todoHandler.AddRoute(httpserverapp.Route{Pattern: "GET /todo/{id}", Handle: getTodoHandler})
	return todoHandler
}
