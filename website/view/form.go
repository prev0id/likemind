package view

import "github.com/a-h/templ"

type Input struct {
	Type        string
	Placeholder string
	Name        string
	Value       string
	Required    bool
	Disabled    bool
	Attributes  templ.Attributes
}

type InputLabel struct {
	Text       string
	For        string
	Attributes templ.Attributes
}

type Button struct {
	Type          string
	ID            string
	Attributes    templ.Attributes
	PopoverTarget string
	PopoverAction string
	Disabled      bool
	Htmx          HTMX
}

type Form struct {
	Htmx       HTMX
	Attributes templ.Attributes
}

type HTMX struct {
	Post     string
	Delete   string
	Target   string
	Swap     string
	Encoding string
	Trigger  string
}

func (htmx HTMX) Attributes() templ.Attributes {
	attributes := make(templ.Attributes)
	if htmx.Post != "" {
		attributes["hx-post"] = htmx.Post
	}
	if htmx.Delete != "" {
		attributes["hx-delete"] = htmx.Delete
	}
	if htmx.Target != "" {
		attributes["hx-target"] = htmx.Target
	}
	if htmx.Swap != "" {
		attributes["hx-swap"] = htmx.Swap
	}
	if htmx.Trigger != "" {
		attributes["hx-trigger"] = htmx.Trigger
	}
	if htmx.Encoding != "" {
		attributes["hx-encoding"] = htmx.Encoding
	}
	return attributes
}

type Details struct {
	Summary string
	Open    bool
}

type TextArea struct {
	Name        string
	Placeholder string
	Rows        int
	Value       string
	Required    bool
}
type Checkbox struct {
	ID      string
	Name    string
	Value   string
	Text    string
	Checked bool
	Htmx    HTMX
}
