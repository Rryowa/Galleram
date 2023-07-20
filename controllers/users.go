package controllers

import (
	"galleryChi/views"
	"net/http"
)

type Users struct {
	Templates
}

type Templates struct {
	New views.Template
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
