package widget_registry

import (
	"encoding/json"
	"fmt"

	"likemind/website/widget/footer"
	"likemind/website/widget/header"
	"likemind/website/widget/login_form"
	"likemind/website/widget/notification"
	"likemind/website/widget/tabs"

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
		return header.Component(header.NewState(header.AuthorizedTabs, header.PeopleTab)), nil
	},
	"login_form": func(data []byte) (templ.Component, error) {
		return login_form.Component(), nil
	},
	"tabs": func(data []byte) (templ.Component, error) {
		if data == nil {
			return tabs.Component(tabs.State{}), nil
		}

		state := &tabs.State{}

		if err := json.Unmarshal(data, state); err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}

		fmt.Println(*state)

		return tabs.Component(*state), nil
	},
}
