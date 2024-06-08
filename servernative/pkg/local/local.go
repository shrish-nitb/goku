package local

import (
	"log"
	"net/http"
	// "strings"
)

type Document struct {
	Name  string
	Value []interface{}
}

type Prefix []string

func (doc *Document) CreateHandlers(prefix Prefix) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
	})
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				doc.read(w, r)
			}
		case "POST":
			{
				doc.create(w, r)
			}
		case "PATCH":
			{
				doc.update(w, r)
			}
		case "DELETE":
			{
				doc.delete(w, r)
			}
		}
	})
	return router
}
