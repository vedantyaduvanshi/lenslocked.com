package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

//NewUsers is used to create a new User controller.
// This function will panic if the templates are not parsed correctly and shou;ld only be used during initial setup!!
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("khudse", "views/users/new.html"),
	}
}

type Users struct {
	NewView *views.View
}

// This is used to render the form where a User cam Create a new User account!!
//Get /Signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

//This is used to process the signup form when user try ti create a new account

//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is fake messgae.Pretend that we created the user account")
}
