package servernativeapp

import (
	"log"
	"servernative/pkg/httpserver/httpserverapp"
	"servernative/pkg/httpserver/httpserverapp/middleware"
	"servernative/pkg/todo"
)

func Run() {
	todos := todo.NewTodos()

	todoHandler := todos.CreateHandlers()

	config := httpserverapp.Config{
		Addr: "127.0.0.1:8000",
	}

	masterHandle := httpserverapp.New()

	//-> Logger <-> Parser <-> todoHandler
	masterHandle.Use(middleware.Logger(masterHandle))
	masterHandle.Use(middleware.BodyParser(masterHandle))
	masterHandle.Pass(todoHandler)
	todoHandler.Pass(masterHandle)
	masterHandle.Use(middleware.Logger(masterHandle))
	log.Println("masterHandle at ", &masterHandle)
	log.Println("todoHandle at ", &todoHandler)

	httpserverapp.Run(&config, masterHandle)
}
