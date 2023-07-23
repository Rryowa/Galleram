package controllers

import (
	"fmt"
	"galleryChi/models"
	"net/http"
)

//Using Template Interface to avoid views dependency

type Users struct {
	UpTpl       Template
	InTpl       Template
	UserService *models.UserService
}

func (u Users) SignUp(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.UpTpl.Execute(w, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.InTpl.Execute(w, data)
}

func (u Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Create(
		&models.NewUser{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Went wrong while creating", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) AuthUser(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Auth(
		&models.NewUser{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "User or password not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User authenticated: %+v", user)
}
