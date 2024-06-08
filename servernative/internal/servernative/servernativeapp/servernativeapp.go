package servernativeapp

import (
	"fmt"
	"servernative/pkg/httpserver/httpserverapp"
	"servernative/pkg/local"
)

func Run() {
	root := httpserverapp.Root{
		Name: "/api/v1",
	}

	todos := local.Document{
		Name: "todo",
	}

	todoHandler := todos.CreateHandlers(local.Prefix{"todo"})

	root.Handlers = append(root.Handlers, todoHandler)

	rootCollection := httpserverapp.RootList{root}

	fmt.Println(rootCollection)

	router := rootCollection.Bind()

	config := httpserverapp.Config{
		Addr: "127.0.0.1:8000",
	}

	httpserverapp.Run(&config, router)
}
