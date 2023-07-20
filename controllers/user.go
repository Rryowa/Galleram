package controllers

import (
	"fmt"
	"net/http"
)

//Using Template Interface to avoid views dependency

type User struct {
	Templates struct {
		Tpl Template
	}
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email tring
	}
	data.Email = r.FormValue("email")
	u.Templates.Tpl.Execute(w, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Email: ", r.PostFormValue("email"))
	fmt.Fprintln(w, "Pass: ", r.PostFormValue("password"))
}
