package app

import (
	"net/http"

	"github.com/iggyster/lets-go-chat/internal/handler"
)

type Router struct {
	routes map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{routes: make(map[string]http.Handler)}
}

func (r *Router) Add(method string, path string, handler http.Handler) {
	r.routes[method+path] = handler
}

func (r *Router) Find(method string, path string) http.Handler {
	h, ok := r.routes[method+path]
	if !ok {
		h = &handler.NotFound{}
	}

	return h
}
