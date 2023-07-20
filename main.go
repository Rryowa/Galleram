package main

import (
	"galleryChi/templates"
	"galleryChi/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, filename string) {
	t := views.Must(views.ParseFS(templates.FS, filename, "tailwind.gohtml"))
	//	Must -> if err != nil {panic}
	//if err != nil {
	//	//log error in terminal
	//	//log.Printf("parsing template: %v", err) (we dont need this if we use chi.Recoverer)
	//	//tell the server about error (we dont need this if we use chi.Recoverer)
	//	//http.Error(w, "There was an error parsing template", http.StatusInternalServerError)
	//	panic(err)
	//}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home.gohtml")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact.gohtml")
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "signup.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Recoverer recover from panic, sends error 500 and prints traced log in terminal
	r.Use(middleware.Recoverer)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/signup", signupHandler)
	r.NotFoundHandler()
	http.ListenAndServe(":3000", r)
}
