package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

//NewUsers Function is used to create a new users controller
// this funtion will be panic if the templates are not parsed correctly.
// Used only for initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

//Users Type
type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can
// create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// comment
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process a the signup form when a user
// submin it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
