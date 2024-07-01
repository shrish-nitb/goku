package middleware

import (
	"log"
	"servernative/pkg/httpserver/httpserverapp"
)

func BodyParser() httpserverapp.HandlerFunc {
	return httpserverapp.HandlerFunc(func(h *httpserverapp.Handle) {
		h.Context.Set("BODY", "BODYPARSED")
		log.Println("Request passed through bodyparser")
		h.Next()
	})
}
