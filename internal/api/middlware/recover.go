package middlware

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if r == http.ErrAbortHandler {
					panic(r)
				}

				slog.Error("recovered go panic")

				if err, ok := r.(error); ok && err != nil {
					slog.Error("panic error", slog.String("error", err.Error()))
				}

				fmt.Printf(string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
