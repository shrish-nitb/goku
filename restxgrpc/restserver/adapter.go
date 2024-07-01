package restserver

import (
	"context"
	"log"
	"net/http"
	pb "restxgrpc/gen/protos/todopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	//GRPC Client
	conn, err := grpc.Dial("127.0.0.1:5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("✅ GRPC Client <-> GRPC Server")
	defer conn.Close()
	client := pb.NewTodoServiceClient(conn)

	context := context.Background()

	//REST Wrapper
	mux := http.NewServeMux()
	mux.HandleFunc("POST /todo", Create(client, context))
	mux.HandleFunc("GET /todo", Read(client, context))
	mux.HandleFunc("PUT /todo", Update(client, context))
	mux.HandleFunc("DELETE /todo", Delete(client, context))

	log.Println("REST Server started at 127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))

}
