package parser

import "net/http"

func Parser(wrappedMiddleware http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedMiddleware.ServeHTTP(w, r)
	})
}
