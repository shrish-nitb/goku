package todo

import (
	"fmt"
	"log"
	"net/http"
	pb "servernativebasic/gen/protos/todopb"
	"time"

	"google.golang.org/protobuf/proto"
)

func (List TodoList) ReadProto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		response, err := proto.Marshal(&pb.TodoListResponse{List: List})

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Println(string(response))

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(response)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, "ns Response Size: ", responseSize, "bytes")
	})
}
