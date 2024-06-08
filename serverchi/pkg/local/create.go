package local

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (list TodoList) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var TodoMessage TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.Trim(string(TodoMessage.Task), " ") == "" {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := list[TodoMessage.Id]; exist {
			http.Error(w, "Resource already exist", http.StatusConflict)
			return
		}
		fmt.Println("qqq")
		list[TodoMessage.Id] = TodoMessage.Task
		fmt.Println("qq")
		response, err := json.Marshal(list)

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}
