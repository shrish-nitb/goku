package restserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	pb "restxgrpc/gen/protos/todopb"
	"time"
)

func Read(client pb.TodoServiceClient, ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		res, err := client.ReadTodo(context.Background(), &pb.Empty{})
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
