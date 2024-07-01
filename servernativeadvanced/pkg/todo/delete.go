package todo

import (
	"net/http"
	"strings"
)

func (List TodoList) Delete(TodoMessage *TodoMessage) (TodoList, *Error) {
	if strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(string(TodoMessage.Id)) == "" {
		return nil, &Error{
			Message: "Invalid request payload",
			Code:    http.StatusBadRequest,
		}
	}

	if _, exist := List[TodoMessage.Id]; !exist {
		return nil, &Error{
			Message: "Resource not exist",
			Code:    http.StatusNotFound,
		}
	}

	List[TodoMessage.Id].Mutex.Lock()
	defer List[TodoMessage.Id].Mutex.Unlock()

	delete(List, TodoMessage.Id)

	return List, nil
}
