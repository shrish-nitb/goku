package local

import (
	"context"
	"net/http"
	"sync"
)

func (List TodoList) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TodoMessage, _ := r.Context().Value(MessageFlagKey).(TodoMessage)
		var ctx context.Context

		if _, exist := List[TodoMessage.Id]; exist {
			http.Error(w, "Resource already exist", http.StatusConflict)
			ctx = context.WithValue(r.Context(), ErrorFlagKey, List)

			req := *r.WithContext(ctx)
			*r = req
			return
		}

		var mu sync.Mutex
		List[TodoMessage.Id] = Task{Mutex: &mu, Value: TodoMessage.Task.Value}

		ctx = context.WithValue(r.Context(), ResponseFlagKey, List)

		*r = *r.WithContext(ctx)
	})
}
