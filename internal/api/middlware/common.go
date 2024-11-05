package middlware

import (
	"net/http"
	"slices"
)

type Middleware func(http.Handler) http.Handler

func Join(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range slices.Backward(middlewares) {
		handler = middleware(handler)
	}

	return handler
}
