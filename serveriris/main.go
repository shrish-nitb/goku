package main

import (
	"serveriris/pkg/todo"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	todos := todo.NewTodos()
	const TODOMESSAGE = "todomessage"

	//parsing middleware
	app.Use(func(ctx iris.Context) {
		todoMessage := new(todo.TodoMessage)
		if err := todoMessage.Parse(ctx); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": "Invalid JSON"})
			return
		}
		ctx.Values().Set(TODOMESSAGE, todoMessage)
		ctx.Next()
	})
	app.Get("/todo", func(ctx iris.Context) {
		todoList, err := todos.Read()
		if err != nil {
			ctx.StatusCode(err.Code)
			ctx.JSON(err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(todoList)
	})
	app.Post("/todo", func(ctx iris.Context) {
		todoList, err := todos.Create(ctx.Values().Get(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			ctx.StatusCode(err.Code)
			ctx.JSON(err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(todoList)
	})
	app.Put("/todo", func(ctx iris.Context) {
		todoList, err := todos.Update(ctx.Values().Get(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			ctx.StatusCode(err.Code)
			ctx.JSON(err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(todoList)
	})
	app.Delete("/todo", func(ctx iris.Context) {
		todoList, err := todos.Delete(ctx.Values().Get(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			ctx.StatusCode(err.Code)
			ctx.JSON(err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(todoList)
	})
	app.Listen(":3000")
}
