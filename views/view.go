package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateExt = "*.gohtml"
)

// Render will render the template with data
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles returns a lits of file paths for the layout templates
func layoutFiles() []string {
	layouts, err := filepath.Glob(LayoutDir + TemplateExt)
	if err != nil {
		panic(err)
	}
	return layouts
}

// NewView will append layouts to data templates
func NewView(layout string, files ...string) *View {

	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View - holds the parsed templates
type View struct {
	Template *template.Template
	Layout   string
}
