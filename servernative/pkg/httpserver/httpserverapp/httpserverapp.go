package httpserverapp

import (
	"log"
	"net/http"
	"servernative/pkg/httpserver/middleware/logger"
	"servernative/pkg/httpserver/middleware/parser"
)

type Config struct {
	Addr string
	Root []Root
}

type Root struct {
	Name     string
	Handlers []http.Handler
}

type RootList []Root

func (RootList *RootList) Bind() *http.ServeMux {
	rootController := http.NewServeMux()
	for _, root := range *RootList {
		for _, handler := range root.Handlers {
			rootController.Handle(root.Name, handler)
		}
	}

	return rootController
}

func Chainer(Router http.Handler, middlewareCollection []func(http.Handler) http.Handler) http.Handler {
	wrappedMiddleware := Router
	for i := len(middlewareCollection); i > 0; i-- {
		wrappedMiddleware = middlewareCollection[i-1](wrappedMiddleware)
	}
	return wrappedMiddleware
}

func Run(config *Config, Router http.Handler) {
	middlewareCollection := []func(http.Handler) http.Handler{logger.Logger, parser.Parser}
	masterHandle := Chainer(Router, middlewareCollection)
	server := http.Server{
		Addr:    config.Addr,
		Handler: masterHandle,
	}
	log.Println("server started")
	server.ListenAndServe()
}
