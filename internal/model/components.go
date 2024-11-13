package model

import (
	"github.com/a-h/templ"
)

type Checkbox struct {
	Label   string
	Name    string
	Checked bool
	Class   string
	Attrs   templ.Attributes
}

type MenuItem struct {
	Name string
	URL  string
	Icon string
}
