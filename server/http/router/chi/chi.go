// Package chi provides a Chi implementation of the router interface.
package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-micro/plugins/server/http/router/router"
)

var _ router.Router = (*Router)(nil)

func init() {
	if err := router.Plugins.Add("chi", newRouter); err != nil {
		panic(err)
	}
}

// Router router implements the router interface.
type Router struct {
	*chi.Mux
}

func newRouter() router.Router {
	return NewRouter()
}

// NewRouter creates a new Chi router.
func NewRouter() *Router {
	r := Router{
		Mux: chi.NewRouter(),
	}

	return &r
}

// Middlewares returns the list of currently registered middlewares.
func (r *Router) Middlewares() []func(http.Handler) http.Handler {
	m := r.Mux.Middlewares()

	return []func(http.Handler) http.Handler(m)
}
