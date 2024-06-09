package local

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func (List TodoList) Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TodoMessage, _ := r.Context().Value("TodoMessage").(TodoMessage)

		if _, exist := List[TodoMessage.Id]; exist {
			http.Error(w, "Resource already exist", http.StatusConflict)
			return
		}

		var mu sync.Mutex
		List[TodoMessage.Id] = Task{Mutex: &mu, Value: TodoMessage.Task.Value}

		fmt.Println(List)
		//Till here it is working as intended

		ctx := context.WithValue(r.Context(), "Response", List)
		r = r.WithContext(ctx)

	})
}
