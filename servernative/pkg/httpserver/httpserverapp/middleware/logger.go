package middleware

import (
	"log"
	"net/http"
	"servernative/pkg/httpserver/httpserverapp"
)

func Logger(handle *httpserverapp.Handle) httpserverapp.HandlerFunc {
	return httpserverapp.HandlerFunc(func(ctx *httpserverapp.Context, w http.ResponseWriter, r *http.Request) {
		log.Printf("Request passed through logger")
		handle.Next()
	})
}
