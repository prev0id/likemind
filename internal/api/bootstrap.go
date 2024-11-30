package api

import (
	dev_handlers "likemind/internal/api/handlers/dev_handler"
	profile_handlers "likemind/internal/api/handlers/profile_handler"
	"likemind/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BootstrapServer(cfg config.API) error {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")

	ph := profile_handlers.New(nil)
	e.GET("/profile", ph.ProfilePage)
	e.GET("/sign_in", ph.SignIn)
	e.GET("/register", ph.Register)
	e.GET("/search/user", ph.SearchUser)
	api.GET("/user/:username", ph.ProfilePage)
	api.POST("/user/:username/update_name", ph.ChangeName)

	dev := e.Group("/dev")

	dh := dev_handlers.New()
	dev.GET("", dh.Page)
	dev.GET("/widget/:widget", dh.MockWidget)
	dev.GET("/widget/list_mocks", dh.ListMocks)

	e.Static("/static", "website/static")

	e.HTTPErrorHandler = ph.Error

	return e.Start(cfg.Addr)
}
