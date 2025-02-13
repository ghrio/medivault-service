package middleware

import "net/http"

type Middleware struct {
	handler http.Handler
}

func ChainMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r)
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{handler: handler}
}
