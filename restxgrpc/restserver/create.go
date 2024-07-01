package restserver

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	pb "restxgrpc/gen/protos/todopb"
	"time"
)

func Create(client pb.TodoServiceClient, ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		TodoMessage := pb.TodoMessageRequest{}
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil || json.Unmarshal(body, &TodoMessage) != nil {
			http.Error(w, "Unmarshal Error", http.StatusBadRequest)
			return
		}
		res, err := client.CreateTodo(ctx, &TodoMessage)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "Marshalling error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, " Response Size: ", responseSize, "bytes")
	})
}
