package serverchiapp

import (
	"net/http"
	"serverchi/pkg/local"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	var todos = local.NewTodos()

	router.Get("/todos", todos.Read())

	router.Post("/todo", todos.Create())

	router.Patch("/todo", todos.Update())

	router.Delete("/todo", todos.Delete())

	http.ListenAndServe(":3000", router)
}
