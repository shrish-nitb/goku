package todo

import (
	"net/http"
	"strings"
	"sync"
)

func (List TodoList) Create(TodoMessage *TodoMessage) (TodoList, *Error) {
	if strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(string(TodoMessage.Id)) == "" {
		return nil, &Error{
			Message: "Invalid request payload",
			Code:    http.StatusBadRequest,
		}
	}

	if _, exist := List[TodoMessage.Id]; exist {
		return nil, &Error{
			Message: "Resource already exist",
			Code:    http.StatusConflict,
		}
	}

	var mu sync.Mutex
	List[TodoMessage.Id] = Task{Mutex: &mu, Value: TodoMessage.Task.Value}

	return List, nil
}
