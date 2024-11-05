package middlware

import (
	"log/slog"
	"net/http"
)

type loggingWriter struct {
	w          http.ResponseWriter
	statusCode int
}

func (l *loggingWriter) Write(data []byte) (int, error) {
	l.statusCode = http.StatusOK
	return l.w.Write(data)
}

func (l *loggingWriter) Header() http.Header {
	return l.w.Header()
}

func (l *loggingWriter) WriteHeader(statusCode int) {
	l.statusCode = statusCode
	l.w.WriteHeader(statusCode)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(
			"request",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
		)

		wrapper := &loggingWriter{w: w}

		next.ServeHTTP(wrapper, r)

		slog.Info(
			"response",
			slog.String("status", http.StatusText(wrapper.statusCode)),
		)
	})
}
