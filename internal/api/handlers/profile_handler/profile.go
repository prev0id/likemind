package profile_handlers

import (
	"net/http"

	"likemind/internal/service/profile"
	"likemind/internal/utils"
	error_page "likemind/website/page/error"
	profile_page "likemind/website/page/profile"
	register_page "likemind/website/page/register"
	signin_page "likemind/website/page/signin"
	user_search_page "likemind/website/page/user_search"

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

func (s *Handler) SignIn(c echo.Context) error {
	return utils.Render(c, http.StatusOK, signin_page.Page())
}

func (s *Handler) Register(c echo.Context) error {
	return utils.Render(c, http.StatusOK, register_page.Page())
}

func (h *Handler) ProfilePage(c echo.Context) error {
	return utils.Render(c, http.StatusOK, profile_page.Page())
}

func (h *Handler) SearchUser(c echo.Context) error {
	return utils.Render(c, http.StatusOK, user_search_page.Page())
}

func (h *Handler) ChangeName(c echo.Context) error {
	return nil
}

func (h *Handler) Error(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if echoErr, ok := err.(*echo.HTTPError); ok {
		code = echoErr.Code
	}

	c.Logger().Error(err)

	utils.Render(c, code, error_page.Page(error_page.State{Code: code}))
}
