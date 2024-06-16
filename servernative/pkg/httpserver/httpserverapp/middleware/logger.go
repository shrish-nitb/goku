package middleware

import (
	"log"
	"net/http"
	"servernative/pkg/httpserver/httpserverapp"
)

func Logger(handle *httpserverapp.Handle) {
	handle.Use(httpserverapp.HandlerFunc(func(ctx *httpserverapp.Context, w http.ResponseWriter, r *http.Request) {
		log.Println("Invoked by ", &handle, ": Request passed through logger")
		handle.Next()
	}))
}
