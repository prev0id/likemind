package widget_registry

import (
	"encoding/json"
	"fmt"

	"likemind/website/widget/footer"
	"likemind/website/widget/header"
	"likemind/website/widget/notification"

	"github.com/a-h/templ"
)

var widgets = map[string]WidgetGenerator{
	"notification": func(data []byte) (templ.Component, error) {
		if data == nil {
			return notification.Component(notification.State{}), nil
		}

		state := &notification.State{}

		if err := json.Unmarshal(data, state); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}

		return notification.Component(*state), nil
	},
	"footer": func(data []byte) (templ.Component, error) {
		return footer.Component(), nil
	},
	"header": func(data []byte) (templ.Component, error) {
		return header.Component(), nil
	},
}
