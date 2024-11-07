package widget_registry

import (
	"encoding/json"
	"fmt"
	"likemind/website/widget"

	"github.com/a-h/templ"
)

var widgets = map[string]WidgetGenerator{
	"Notification": func(data []byte) (templ.Component, error) {
		state := &widget.NotificationData{}

		if err := json.Unmarshal(data, state); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}

		return widget.Notification(*state), nil
	},
	"Footer": func(data []byte) (templ.Component, error) {
		return nil, nil
	},
	"Header": func(data []byte) (templ.Component, error) {
		return nil, nil
	},
}
