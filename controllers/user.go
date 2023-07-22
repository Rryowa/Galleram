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

func (u User) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Create(
		&models.NewUser{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("email"),
		})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Went wrong while creating", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "User created: %+v", user)
}
