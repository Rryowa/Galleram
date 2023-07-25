package main

import (
	"galleram/controllers"
	"galleram/models"
	"galleram/templates"
	"galleram/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	db, err := models.Open(models.DefaultPostgresConfig())
	if err != nil {
		panic(err)
	}
	usersController := controllers.Users{
		UpTpl:       views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml")),
		InTpl:       views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml")),
		UserService: &models.UserService{DB: db},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Recoverer recover from panic, sends error 500 and prints traced log in terminal
	r.Use(middleware.Recoverer)
	r.Get("/", controllers.StaticTemplate(views.Must(views.ParseFS(
		templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticTemplate(views.Must(views.ParseFS(
		templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/sign-up", usersController.SignUp)
	r.Get("/sign-in", usersController.SignIn)
	r.Post("/users/new", usersController.CreateUser)
	r.Post("/users/old", usersController.AuthUser)
	r.NotFoundHandler()
	http.ListenAndServe("localhost:8080", r)
}
