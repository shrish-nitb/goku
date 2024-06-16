package main

import (
	"log"
	"net/http"
	"servernativebasic/pkg/todo"
)

func main() {

	todoList := todo.NewTodos()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /proto/todo", todoList.ReadProto())
	mux.HandleFunc("GET /todo", todoList.Read())
	mux.HandleFunc("POST /todo", todoList.Create())
	mux.HandleFunc("PUT /todo", todoList.Update())
	mux.HandleFunc("DELETE /todo", todoList.Delete())

	log.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
