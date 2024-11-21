package dev_handlers

import (
	"fmt"
	"net/http"

	"likemind/internal/service/widget_registry"
	"likemind/website/page"
	"likemind/website/widget/select_dropdown"
)

const (
	defaultMockOption   = "--- Mock ---"
	defaultWidgetOption = "--- Widget ---"
)

type Service struct {
	registry *widget_registry.Registry
}

func New() *Service {
	return &Service{
		registry: widget_registry.New(),
	}
}

func (s *Service) Page(w http.ResponseWriter, r *http.Request) {
	widgets := s.listOfWidgets()

	mocks := select_dropdown.State{
		Label:   "Select a mock",
		ID:      "mock_selection",
		Name:    "mock",
		Default: defaultMockOption,
	}

	page.DevPage(widgets, mocks, nil).Render(r.Context(), w)
}

func (s *Service) MockWidget(w http.ResponseWriter, r *http.Request) {
	widget := r.PathValue("widget")
	if widget == defaultWidgetOption {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	test := query.Get("test")
	if test == defaultMockOption {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	component, err := s.registry.GetWidget(widget, test)
	if err != nil {
		panic(err)
	}

	component.Render(r.Context(), w)
}

func (s *Service) ListMocks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	widget := query.Get("widget")
	if widget == defaultWidgetOption {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	state, err := s.listOfMocks(widget)
	if err != nil {
		panic(err)
	}

	select_dropdown.Component(state).Render(r.Context(), w)
}

func (s *Service) listOfMocks(widgetName string) (select_dropdown.State, error) {
	tests, err := s.registry.ListTestsForWidget(widgetName)
	if err != nil {
		return select_dropdown.State{}, fmt.Errorf("s.registry.ListTestsForWidget widget='%s': %w", widgetName, err)
	}

	return select_dropdown.State{
		Label:   "Select a test",
		ID:      "test_selection",
		Data:    tests,
		Name:    "test",
		Default: defaultMockOption,
		HTMX: &select_dropdown.HTMX{
			Get:    "/dev/mock/" + widgetName,
			Target: "#resizable_wrapper",
		},
	}, nil
}

func (s *Service) listOfWidgets() select_dropdown.State {
	widgets := s.registry.ListWidgets()

	return select_dropdown.State{
		Name:    "widget",
		Label:   "Select a widget",
		ID:      "widget_selection",
		Data:    widgets,
		Default: defaultWidgetOption,
		HTMX: &select_dropdown.HTMX{
			Get:    "/dev/mock",
			Target: "#test_selection_wrapper",
		},
	}
}
