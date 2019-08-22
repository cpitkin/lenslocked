package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is where the global layouts are stored
	LayoutDir = "views/layouts/"

	// TemplateDir is the directory where all the templates live
	TemplateDir = "views/"

	// TemplateExt is the extension used by the layouts
	TemplateExt = ".gohtml"
)

// View - holds the parsed templates
type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := v.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// Render will render the template with data
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

// layoutFiles returns a lits of file paths for the layout templates
func layoutFiles() []string {
	layouts, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return layouts
}

// NewView will append layouts to data templates
func NewView(layout string, files ...string) *View {

	addTemplatePath(files)
	addTemplateExt(files)
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
