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

type HandlerFunc func(ctx *Context, w http.ResponseWriter, r *http.Request)

type Handle struct {
	ctx  *Context
	w    http.ResponseWriter
	r    *http.Request
	pass []HandlerFunc
}

func (h *Handle) init() http.Handler {
	ctx := Context{context.Background()}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ctx = &ctx
		h.w = w
		h.r = r
		h.Next()
	})
}

func (h *Handle) Use(f HandlerFunc) {
	h.pass = append(h.pass, f)
}

func (h *Handle) Pass(successor *Handle) {
	h.Use(HandlerFunc(func(ctx *Context, w http.ResponseWriter, r *http.Request) {
		successor.ctx = h.ctx
		successor.w = h.w
		successor.r = h.r
		successor.Next()
	}))

}

func (h *Handle) Next() {
	if len(h.pass) == 0 {
		return
	}
	tmp := h.pass[0]
	h.pass = h.pass[1:]
	tmp(h.ctx, h.w, h.r)
}

func New() *Handle {
	return &Handle{}
}

func Run(config *Config, masterHandle *Handle) {
	server := http.Server{
		Addr:    config.Addr,
		Handler: masterHandle.init(),
	}
	log.Printf("server started at %s", config.Addr)
	server.ListenAndServe()
}
