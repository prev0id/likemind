package dev_handlers

// import (
// 	"fmt"
// 	"net/http"

// 	"likemind/internal/service/widget_registry"
// 	dev_page "likemind/website/page/dev"
// 	"likemind/website/widget/select_dropdown"

// 	"github.com/a-h/templ"
// 	"github.com/labstack/echo/v4"
// )

// const (
// 	mockParam   = "mock"
// 	widgetParam = "widget"
// )

// type mockService interface {
// 	ListWidgets() []string
// 	ListWidgetMocks(widget string) ([]string, error)
// 	GetWidgetMock(widget, mock string) (templ.Component, error)
// }

// type Service struct {
// 	svc mockService
// }

// func New() *Service {
// 	return &Service{
// 		svc: widget_registry.New(),
// 	}
// }

// func (s *Service) Page(c echo.Context) error {
// 	widgets := s.svc.ListWidgets()

// 	defaultWidgetName := widgets[0]
// 	mocks, err := s.svc.ListWidgetMocks(defaultWidgetName)
// 	if err != nil {
// 		return err
// 	}

// 	defaultMockName := mocks[0]
// 	defaultMock, err := s.svc.GetWidgetMock(defaultWidgetName, defaultMockName)
// 	if err != nil {
// 		return err
// 	}

// 	devPage := dev_page.Page(widgetSelector(widgets), mockSelector(mocks, defaultWidgetName), defaultMock)

// 	return utils.Render(c, http.StatusOK, devPage)
// }

// func (s *Service) ListMocks(c echo.Context) error {
// 	widget := c.QueryParam(widgetParam)

// 	mocks, err := s.svc.ListWidgetMocks(widget)
// 	if err != nil {
// 		return utils.RenderErrorNotification(c, http.StatusBadRequest, err)
// 	}

// 	return utils.Render(c, http.StatusOK, mockSelector(mocks, widget))
// }

// func (s *Service) MockWidget(c echo.Context) error {
// 	widget := c.Param(widgetParam)
// 	mock := c.QueryParam(mockParam)

// 	fmt.Println(widget, mock)

// 	component, err := s.svc.GetWidgetMock(widget, mock)
// 	if err != nil {
// 		return utils.RenderErrorNotification(c, http.StatusBadRequest, err)
// 	}

// 	return utils.Render(c, http.StatusOK, component)
// }

// func mockSelector(mocks []string, widgetName string) templ.Component {
// 	return select_dropdown.Component(
// 		select_dropdown.State{
// 			Label:      "Select a mock",
// 			ID:         "mock_selection",
// 			Data:       mocks,
// 			Name:       mockParam,
// 			HTMXGet:    "/dev/widget/" + widgetName,
// 			HTMXTarget: "#resizable_wrapper",
// 		},
// 	)
// }

// func widgetSelector(widgets []string) templ.Component {
// 	return select_dropdown.Component(
// 		select_dropdown.State{
// 			Name:       widgetParam,
// 			Label:      "Select a widget",
// 			ID:         "widget_selection",
// 			Data:       widgets,
// 			HTMXGet:    "/dev/widget/list_mocks",
// 			HTMXTarget: "#mock_selection_wrapper",
// 		},
// 	)
// }
