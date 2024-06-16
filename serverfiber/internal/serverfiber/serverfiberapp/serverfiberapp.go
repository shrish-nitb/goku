package serverfiberapp

import (
	"serverfiber/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()
	todos := todo.NewTodos()
	const TODOMESSAGE = "todomessage"

	//parsing middleware
	app.Use(func(ctx *fiber.Ctx) error {
		todoMessage := new(todo.TodoMessage)
		if err := todoMessage.Parse(ctx); err != nil {
			return ctx.Status(fiber.ErrBadRequest.Code).JSON(todo.Error{Message: "Bad Request", Code: fiber.ErrBadRequest.Code})
		}
		ctx.Locals(TODOMESSAGE, todoMessage)
		return ctx.Next()
	})

	app.Get("/todo", func(ctx *fiber.Ctx) error {
		todoList, err := todos.Read()
		if err != nil {
			return ctx.Status(err.Code).JSON(err)
		}
		return ctx.Status(fiber.StatusOK).JSON(todoList)
	})
	app.Post("/todo", func(ctx *fiber.Ctx) error {
		todoList, err := todos.Create(ctx.Locals(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			return ctx.Status(err.Code).JSON(err)
		}
		return ctx.Status(fiber.StatusOK).JSON(todoList)
	})
	app.Put("/todo", func(ctx *fiber.Ctx) error {
		todoList, err := todos.Update(ctx.Locals(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			return ctx.Status(err.Code).JSON(err)
		}
		return ctx.Status(fiber.StatusOK).JSON(todoList)
	})
	app.Delete("/todo", func(ctx *fiber.Ctx) error {
		todoList, err := todos.Delete(ctx.Locals(TODOMESSAGE).(*todo.TodoMessage))
		if err != nil {
			return ctx.Status(err.Code).JSON(err)
		}
		return ctx.Status(fiber.StatusOK).JSON(todoList)
	})
	app.Listen(":3000")
}
