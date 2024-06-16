package httpserverapp

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Addr string
}

type Context struct {
	context.Context
}

func (c *Context) Set(key any, value any) {
	*c = Context{context.WithValue(*c, key, value)}
}

func (c *Context) Get(key any) any {
	return c.Value(key)
}

type HandlerFunc func(h *Handle)
type Pattern struct {
	Target string
	Handle *Handle
}
type Handle struct {
	Context *Context
	Writer  http.ResponseWriter
	Request *http.Request
	mux     *http.ServeMux
	pass    []HandlerFunc
	routes  []Pattern
}

func (h *Handle) init() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := h.Context
		if ctx == nil {
			ctx = &Context{context.Background()}
		}
		newHandle := h.clone()
		fmt.Printf("New Handler Cloned at %p\n", newHandle)
		newHandle.Context = ctx
		newHandle.Writer = w
		newHandle.Request = r
		newHandle.Use(HandlerFunc(func(hNew *Handle) {
			fmt.Println("Request came to Routing Middleware")
			for _, pattern := range h.routes {
				hNew.mux.Handle(pattern.Target, http.StripPrefix("/todo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					pattern.Handle.Context = hNew.Context
					pattern.Handle.init().ServeHTTP(w, r)
				})))
			}
			hNew.mux.ServeHTTP(hNew.Writer, hNew.Request)
		}))
		newHandle.Next()
	})
}

func (h *Handle) clone() *Handle {
	clone := New()
	clone.pass = append(clone.pass, h.pass...)
	clone.routes = append(clone.routes, h.routes...)
	return clone
}

func (h *Handle) AddRouter(pattern Pattern) {
	h.routes = append(h.routes, pattern)
}

func (h *Handle) Pass(successor *Handle) {
	h.Use(HandlerFunc(func(h *Handle) {
		fmt.Printf("Request passed through Handle Transfer Middleware shutting %p\n", h)
		successor.Context = h.Context
		successor.init().ServeHTTP(h.Writer, h.Request)
	}))
}

func (h *Handle) Use(f HandlerFunc) {
	h.pass = append(h.pass, f)
}

func (h *Handle) Next() {
	fmt.Printf("Invoked by %p\n", h)
	tmp := h.pass[0]
	h.pass = h.pass[1:]
	tmp(h)
}

func New() *Handle {
	return &Handle{mux: http.NewServeMux()}
}

func Run(config *Config, masterHandle *Handle) {
	server := http.Server{
		Addr:    config.Addr,
		Handler: masterHandle.init(),
	}
	log.Printf("server started at %s", config.Addr)
	server.ListenAndServe()
}
