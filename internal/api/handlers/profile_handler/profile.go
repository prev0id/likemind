package profile_handlers

import (
	"net/http"

	"likemind/internal/service/profile"
	"likemind/internal/utils"
	"likemind/website/page"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	profileService profile.Service
}

func New(profileService profile.Service) *Handler {
	return &Handler{
		profileService: profileService,
	}
}

func (s *Handler) Login(c echo.Context) error {
	return nil
}

func (s *Handler) Register(c echo.Context) error {
	return utils.Render(c, http.StatusOK, page.Regeister())
}

func (h *Handler) ProfilePage(c echo.Context) error {
	return utils.Render(c, http.StatusOK, page.ProfilePage())
}

func (h *Handler) ChangeName(c echo.Context) error {
	return nil
}

func RegistrationPage(w http.ResponseWriter, r *http.Request) {
}
