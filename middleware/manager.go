package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) With(middlwares ...Middleware) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		n := next

		// [hudai(logger(next)), ]
		for i := len(middlwares) - 1; i >= 0; i-- {
			middlware := middlwares[i]
			n = middlware(n)
		}

		return n
	}
}
