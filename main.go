package main

import (
	"galleryChi/controllers"
	"galleryChi/templates"
	"galleryChi/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Recoverer recover from panic, sends error 500 and prints traced log in terminal
	r.Use(middleware.Recoverer)
	//r.Get("/", homeHandler)
	//r.Get("/contact", contactHandler)
	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml",
		"tailwind.gohtml",
	))
	r.Get("/signup", usersC.New)
	r.NotFoundHandler()
	http.ListenAndServe(":3000", r)
}
