package middleware_custom

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"serverchi/pkg/local"
	"strings"
)

func Parser(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var TodoMessage local.TodoMessage
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil || json.Unmarshal(body, &TodoMessage) != nil || strings.Trim(string(TodoMessage.Task.Value), " ") == "" {
			http.Error(w, "ParseErr: Invalid request payload", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "TodoMessage", TodoMessage)
		newReq := r.WithContext(ctx)

		f.ServeHTTP(w, newReq)

		List := newReq.Context().Value("Response")

		response, err := json.Marshal(List)

		if err != nil {
			http.Error(w, "ParseErr: Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})
}
