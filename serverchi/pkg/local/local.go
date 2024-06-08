package local

type Id string
type Task string

type TodoMessage struct {
	Id   Id   `json:"id"`
	Task Task `json:"task"`
}

type TodoList map[Id]Task

func NewTodos() *TodoList {
	return new(TodoList)
}
