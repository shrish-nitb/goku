package todo

import (
	"sync"

	"github.com/gofiber/fiber/v2"
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

func (todoMessage *TodoMessage) Parse(c *fiber.Ctx) error {
	if err := c.BodyParser(todoMessage); err != nil {
		return err
	}
	return nil
}
