package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Manager struct {
	globalMiddlwares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlwares: make([]Middleware, 0),
	}
}

func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		n := next

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			n = middleware(n)
		}

		return n;
 	}
	
}