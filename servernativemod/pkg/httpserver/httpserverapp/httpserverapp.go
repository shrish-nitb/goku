package httpserverapp

import (
	"context"
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
type Route struct {
	Pattern string
	Handle  *Handle
}
type Handle struct {
	Context     *Context
	Writer      http.ResponseWriter
	Request     *http.Request
	mux         *http.ServeMux
	middlewares []HandlerFunc
	routes      []Route
}

func (h *Handle) init() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := h.Context
		if ctx == nil {
			ctx = &Context{context.Background()}
		}
		newHandle := h.clone()
		log.Printf("New Handle Cloned at %p\n", newHandle)
		newHandle.Context = ctx
		newHandle.Writer = w
		newHandle.Request = r
		newHandle.Use(HandlerFunc(func(hNew *Handle) {
			log.Println("Request came to Routing Middleware")
			for _, route := range h.routes {
				hNew.mux.Handle(route.Pattern, http.StripPrefix("/todo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					route.Handle.Context = hNew.Context
					route.Handle.init().ServeHTTP(w, r)
				})))
			}
			hNew.mux.ServeHTTP(hNew.Writer, hNew.Request)
		}))
		newHandle.Next()
	})
}

func (h *Handle) clone() *Handle {
	clone := New()
	clone.middlewares = append(clone.middlewares, h.middlewares...)
	clone.routes = append(clone.routes, h.routes...)
	return clone
}

func (h *Handle) AddRoute(route Route) {
	h.routes = append(h.routes, route)
}

func (h *Handle) Pass(successor *Handle) {
	h.Use(HandlerFunc(func(h *Handle) {
		log.Printf("Request passed through Handle Transfer Middleware shutting %p\n", h)
		successor.Context = h.Context
		successor.init().ServeHTTP(h.Writer, h.Request)
	}))
}

func (h *Handle) Use(f HandlerFunc) {
	h.middlewares = append(h.middlewares, f)
}

func (h *Handle) Next() {
	log.Printf("Invoked by %p\n", h)
	tmp := h.middlewares[0]
	h.middlewares = h.middlewares[1:]
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
