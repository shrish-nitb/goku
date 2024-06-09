package local

import (
	"encoding/json"
	"io"
	"net/http"
)

func (List TodoList) ReadOne() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var TodoMessage TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil || json.Unmarshal(body, &TodoMessage) != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, exist := List[TodoMessage.Id]; !exist {
			http.Error(w, "Resource Not Found", http.StatusNotFound)
			return
		}

		response, err := json.Marshal(List[TodoMessage.Id])

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}
