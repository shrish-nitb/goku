package todo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func (List TodoList) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		var TodoMessage TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(string(TodoMessage.Id)) == "" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := List[TodoMessage.Id]; exist {
			http.Error(w, "Resource already exist", http.StatusConflict)
			return
		}

		var mu sync.Mutex
		List[TodoMessage.Id] = Task{Mutex: &mu, Value: TodoMessage.Task.Value}

		response, err := json.Marshal(List)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, " Response Size: ", responseSize, "bytes")
	})
}
