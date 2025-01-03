package page_handler

import (
	"log"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"

	profile_page "likemind/website/page/profile"
	register_page "likemind/website/page/register"
	signin_page "likemind/website/page/signin"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type PageHandler struct{}

func New() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) Prefix() string {
	return domain.PathPrefixRoot
}

func (h *PageHandler) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get(domain.PathUserPage, common.WrapHTMLHandler(h.profile))
	router.Get(domain.PatternSignIn, common.WrapHTMLHandler(h.sigin))
	router.Get(domain.PatternSignUp, common.WrapHTMLHandler(h.signup))

	return router
}

func (h *PageHandler) profile(_ http.ResponseWriter, r *http.Request) (component templ.Component, statusCode int) {
	username := chi.URLParam(r, domain.PathVarUsername)
	log.Println(username)
	return profile_page.Page(), http.StatusOK
}

func (h *PageHandler) sigin(_ http.ResponseWriter, _ *http.Request) (component templ.Component, statusCode int) {
	return signin_page.Page(), http.StatusOK
}

func (h *PageHandler) signup(_ http.ResponseWriter, _ *http.Request) (component templ.Component, statusCode int) {
	return register_page.Page(), http.StatusOK
}
