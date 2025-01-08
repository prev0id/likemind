package page_handler

import (
	"log"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/internal/service/auth"

	profile_page "likemind/website/page/profile"
	signin_page "likemind/website/page/signin"
	signup_page "likemind/website/page/signup"
	user_search_page "likemind/website/page/user_search"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type PageHandler struct {
	authSvc auth.Service
}

func New(authSvc auth.Service) *PageHandler {
	return &PageHandler{
		authSvc: authSvc,
	}
}

func (h *PageHandler) Prefix() string {
	return domain.PathPrefixRoot
}

func (h *PageHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Group(func(public chi.Router) {
		router.Get(domain.PathSignIn, common.WrapHTMLHandler(h.signin))
		router.Get(domain.PathSignUp, common.WrapHTMLHandler(h.signup))
	})

	router.Group(func(protected chi.Router) {
		protected.Use(h.authSvc.Middleware)

		protected.Get(domain.PathUserPage, common.WrapHTMLHandler(h.profile))
	})

	return router
}

func (h *PageHandler) profile(_ http.ResponseWriter, r *http.Request) (component templ.Component, statusCode int) {
	username := chi.URLParam(r, domain.PathVarNickname)
	log.Println(username)
	return profile_page.Page(), http.StatusOK
}

func (h *PageHandler) search(_ http.ResponseWriter, _ *http.Request) (component templ.Component, statusCode int) {
	return user_search_page.Page(), http.StatusOK
}

func (h *PageHandler) signin(_ http.ResponseWriter, _ *http.Request) (component templ.Component, statusCode int) {
	return signin_page.Page(), http.StatusOK
}

func (h *PageHandler) signup(_ http.ResponseWriter, _ *http.Request) (component templ.Component, statusCode int) {
	return signup_page.Page(), http.StatusOK
}
