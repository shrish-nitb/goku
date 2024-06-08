package local

import (
	"net/http"
)

func (doc *Document) read(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello"))
	return nil
}
