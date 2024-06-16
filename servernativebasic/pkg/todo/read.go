package todo

import (
	"encoding/json"
	"log"
	"net/http"
	pb "servernativebasic/gen/protos/todopb"
	"time"
)

func (List TodoList) Read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		response, err := json.Marshal(pb.TodoListResponse{List: List})

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
