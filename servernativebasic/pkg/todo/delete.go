package todo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func (List TodoList) Delete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		var TodoMessage TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(string(TodoMessage.Id)) == "" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := List[TodoMessage.Id]; !exist {
			http.Error(w, "Resource Not Found", http.StatusNotFound)
			return
		}

		List[TodoMessage.Id].Mutex.Lock()
		defer List[TodoMessage.Id].Mutex.Unlock()

		delete(List, TodoMessage.Id)

		response, err := json.Marshal(List)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, "ns Response Size: ", responseSize, "bytes")
	})
}
