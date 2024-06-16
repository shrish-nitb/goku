package middleware

import (
	"log"
	"net/http"
	"servernative/pkg/httpserver/httpserverapp"
)

func BodyParser(handle *httpserverapp.Handle) httpserverapp.HandlerFunc {
	return httpserverapp.HandlerFunc(func(ctx *httpserverapp.Context, w http.ResponseWriter, r *http.Request) {
		log.Println("Invoked by ", &handle, "Request passed through bodyparser")
		handle.Next()
	})
}
