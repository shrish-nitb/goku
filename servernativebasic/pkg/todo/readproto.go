package todo

import (
	"log"
	"net/http"
	"servernativebasic/gen/protos/todopb"
	"time"

	"google.golang.org/protobuf/proto"
)

func (List TodoList) ReadProto() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		TodoListResponse := &todopb.TodoListResponse{
			TodoList: ,
		}
		protoResponse, err := proto.Marshal(TodoListResponse)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseSize, _ := w.Write(protoResponse)

		elapsedTime := time.Since(startTime)
		log.Println("Request Time Taken:", elapsedTime, "ns Response Size: ", responseSize, "bytes")
	})
}
