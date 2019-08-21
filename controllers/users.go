package controllers

import (
	"fmt"
	"lenslocked/views"
	"net/http"
)

// Users holds the users controller information
type Users struct {
	NewView *views.View
}

// SignupForm holds the new user signup form data
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUsers creates a new users signup view. This function will panic so it should only be used during set up
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New is used to render the form where a user can create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process a new user creation
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	err := ParseForm(r, &form)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}
