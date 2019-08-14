package main

import (
	"fmt"
	"lenslocked/views"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>Please direct all questions to our email that can be found in the <a href=/contact>contact</a> page.</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you "+"were looking for :(</h1>"+"<p>Please email us if you keep being sent to an invalid page.</p>")
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")

	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	var nf http.Handler = http.HandlerFunc(notFound)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = nf
	http.ListenAndServe(":3000", r)
}
