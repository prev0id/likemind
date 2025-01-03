package common

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
)

var errNoRequest = errors.New("unnable to get request from context")

type (
	Handler      func(w http.ResponseWriter, r *http.Request) (component templ.Component, status int)
	ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)
)

// func BindPage[T any](router chi.Router, method, path string, handler Handler) {
// 	bind[T](router, method, path, handler, errorPage)
// }

// func BindAPI[T any](router chi.Router, method, path string, handler Handler) {
// 	bind[T](router, method, path, handler, errorAPI)
// }

// func bind[T any](router chi.Router, method, path string, handler Handler, errorHandler ErrorHandler) {
// 	inputStruct := new(T)
// 	router.With(
// 		httpin.NewInput(
// 			inputStruct,
// 			httpin.Option.WithErrorHandler(errorHandler),
// 		),
// 	).Method(method, path, render(handler))
// }

func WrapHTMLHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component, statusCode := handler(w, r)

		templ.Handler(component, templ.WithStatus(statusCode)).ServeHTTP(w, r)
	}
}

// func errorPage(w http.ResponseWriter, r *http.Request, _ error) {
// 	handler := func(w http.ResponseWriter, r *http.Request) (int, templ.Component) {
// 		statusCode := http.StatusBadRequest
// 		return statusCode, error_page.Page(error_page.State{Code: statusCode})
// 	}

// 	render(handler)(w, r)
// }

// func errorAPI(w http.ResponseWriter, r *http.Request, err error) {
// 	handler := func(w http.ResponseWriter, r *http.Request) (int, templ.Component) {
// 		statusCode := http.StatusBadRequest
// 		return statusCode, notification.Component(notification.State{
// 			Type:    notification.Error,
// 			Message: err.Error(),
// 		})
// 	}

// 	render(handler)(w, r)
// }
