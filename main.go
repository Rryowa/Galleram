package main

import (
	"galleryChi/controllers"
	"galleryChi/models"
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
	r.Get("/", controllers.StaticTemplate(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticTemplate(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml"))))

	db, err := models.Open(models.DefaultPostgresConfig())
	if err != nil {
		panic(err)
	}
	userService := models.UserService{
		DB: db,
	}
	usersController := controllers.User{
		UserService: &userService,
	}
	usersController.Templates.Tpl = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	r.Get("/users/new", usersController.New)
	r.Post("/users", usersController.Create)
	//submit form
	r.NotFoundHandler()
	http.ListenAndServe("localhost:8080", r)
}
