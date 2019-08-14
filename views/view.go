package views

import "html/template"

// NewView - append layouts to data templates
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// View - holds the parsed templates
type View struct {
	Template *template.Template
}
