package profile_handlers

import (
	"net/http"

	"likemind/internal/service/profile"
	"likemind/website/page"

	"github.com/a-h/templ"
)

type Handler struct {
	profileService profile.Service
}

func New(profileService profile.Service) *Handler {
	return &Handler{
		profileService: profileService,
	}
}

func (h *Handler) ProfilePage(w http.ResponseWriter, r *http.Request) {
	page := page.ProfilePage()

	handler := templ.Handler(page)

	handler.ServeHTTP(w, r)
}

func (h *Handler) ChangeName(w http.ResponseWriter, r *http.Request) {
}

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
}
