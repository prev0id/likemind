package widget_dev

import (
	"likemind/internal/service/widget_registry"
	"likemind/website/page"
	"likemind/website/widget"
	"net/http"
	"slices"

	"github.com/samber/lo"
)

type Service struct {
	registry *widget_registry.Registry
}

func New() *Service {
	return &Service{
		registry: widget_registry.New(),
	}
}

func (s *Service) HandleTestingPage(w http.ResponseWriter, r *http.Request) {
	widgets := s.registry.ListWidgets()

	slices.Sort(widgets)

	data := widget.SelectData{
		Label: "Select Widget",
		ID:    "selection_data",
		Data: lo.Map(widgets, func(w string, _ int) widget.SelectOption {
			return widget.SelectOption{
				Value: w,
				Name:  w,
			}
		}),
	}

	page.TestPage(widget.SelectDropdown(data)).Render(r.Context(), w)
}

func (s *Service) HandleListOfWidgets(w http.ResponseWriter, r *http.Request) {
	widgets := s.registry.ListWidgets()

	data := widget.SelectData{
		Label: "label",
		ID:    "selection_data",
		Data: lo.Map(widgets, func(w string, _ int) widget.SelectOption {
			return widget.SelectOption{
				Value: w,
				Name:  w,
			}
		}),
	}

	widget.SelectDropdown(data).Render(r.Context(), w)
}
