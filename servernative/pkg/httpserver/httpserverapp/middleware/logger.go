package middleware

import (
	"log"
	"servernative/pkg/httpserver/httpserverapp"
)

func Logger() httpserverapp.HandlerFunc {
	return httpserverapp.HandlerFunc(func(h *httpserverapp.Handle) {
		log.Println("Request passed through logger")
		h.Next()
	})
}
