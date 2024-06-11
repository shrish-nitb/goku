package local

import (
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

type contextKey string

const (
	ResponseFlagKey contextKey = "Response"
	ErrorFlagKey    contextKey = "ErrorFlag"
	MessageFlagKey  contextKey = "Message"
)
