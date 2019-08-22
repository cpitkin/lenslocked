package controllers

import (
	"lenslocked/views"
)

// Static holds all the static views
type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}

// NewStatic renders all static templates into a struct
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq:     views.NewView("material", "static/faq"),
	}
}
