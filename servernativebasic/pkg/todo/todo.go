package todo

import (
	pb "servernativebasic/gen/protos/todopb"
)

type TodoList map[string]*pb.Task

func NewTodos() *TodoList {
	var init = make(TodoList)
	return &init
}
