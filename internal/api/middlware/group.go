package middlware

import (
	"net/http"
	"path"
)

type Middleware func(http.Handler) http.Handler

type Group struct {
	prefix      string
	middlewares []Middleware
}

func NewGroup(prefix string, middleware ...Middleware) Group {
	return Group{
		prefix:      prefix,
		middlewares: middleware,
	}
}

func (g Group) Register(method string, pathSuffix string, handler http.Handler) {
	fullPath := path.Join(g.prefix, pathSuffix)

	http.Handle(method+" "+fullPath, g.join(handler))
}

func (g Group) RegisterHandler(method string, pathSuffix string, handler http.Handler) {
	fullPath := path.Join(g.prefix, pathSuffix)

	http.Handle(method+" "+fullPath, g.join(handler))
}

func (g Group) join(handler http.Handler) http.Handler {
	for _, middleware := range g.middlewares {
		handler = middleware(handler)
	}

	return handler
}
