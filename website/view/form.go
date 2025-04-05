package view

import "github.com/a-h/templ"

type Input struct {
	Type        string
	Placeholder string
	Name        string
	Value       string
	Required    bool
	Attributes  templ.Attributes
}

type InputLabel struct {
	Text       string
	For        string
	Attributes templ.Attributes
}

type Button struct {
	Type          string
	Attributes    templ.Attributes
	PopoverTarget string
	PopoverAction string
}

type Form struct {
	Htmx       HTMX
	Attributes templ.Attributes
}

type HTMX struct {
	Post     string
	Target   string
	Swap     string
	Encoding string
}
