package local

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (List TodoList) Update() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var TodoMessage TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.Trim(string(TodoMessage.Task.Value), " ") == "" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := List[TodoMessage.Id]; !exist {
			http.Error(w, "Resource Not Found", http.StatusNotFound)
			return
		}

		List[TodoMessage.Id].Mutex.Lock()
		defer List[TodoMessage.Id].Mutex.Unlock()

		List[TodoMessage.Id] = Task{Mutex: List[TodoMessage.Id].Mutex, Value: TodoMessage.Task.Value}

		response, err := json.Marshal(List)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}