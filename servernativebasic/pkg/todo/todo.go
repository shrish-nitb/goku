package todo

import (
	"sync"
)

type Task struct {
	Mutex *sync.Mutex `json:"-"`
	Value string      `json:"value"`
}

type TodoMessage struct {
	Id   string `json:"id"`
	Task Task   `json:"task"`
}

type TodoList map[string]Task

func NewTodos() *TodoList {
	var init = make(TodoList)
	return &init
}
