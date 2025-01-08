package common

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"likemind/internal/domain"

	"github.com/a-h/templ"
)

var errNoRequest = errors.New("unnable to get request from context")

type (
	Handler      func(w http.ResponseWriter, r *http.Request) (component templ.Component, status int)
	ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)
)

func WrapHTMLHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component, statusCode := handler(w, r)

		templ.Handler(component, templ.WithStatus(statusCode)).ServeHTTP(w, r)
	}
}

func Redirect(w http.ResponseWriter, url string) {
	w.WriteHeader(http.StatusFound)
	w.Header().Set(domain.HTMXRedirectHeader, url)
}

func ServeError(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

var varRegex = regexp.MustCompile(`{\w+}`)

type PathVars map[string]string

func SetPathVariables(path string, values PathVars) string {
	return varRegex.ReplaceAllStringFunc(path, func(pathVar string) string {
		key := strings.Trim(pathVar, "{}")
		val, ok := values[key]
		if !ok {
			return pathVar
		}
		return val
	})
}
