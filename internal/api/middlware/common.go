package middlware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Join(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
