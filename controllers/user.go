package controllers

import (
	"fmt"
	"galleryChi/models"
	"net/http"
)

//Using Template Interface to avoid views dependency

type User struct {
	Templates struct {
		Tpl Template
	}
	UserService *models.UserService
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.Tpl.Execute(w, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	//TODO: HOW TO PASS A STRUCT HERE?
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println()
		http.Error(w, "Went wrong while creating", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "User created: %+v", user)
}
