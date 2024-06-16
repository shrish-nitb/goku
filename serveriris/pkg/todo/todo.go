package todo

import (
	"sync"

	"github.com/kataras/iris/v12"
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

func (todoMessage *TodoMessage) Parse(ctx iris.Context) error {
	if err := ctx.ReadJSON(todoMessage); err != nil {
		return err
	}
	return nil
}
