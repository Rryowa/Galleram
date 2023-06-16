package main

import (
	"galleryChi/views"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, filename string) {
	t, err := views.ParseTemplate(filename)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home.gohtml")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact.gohtml")
}

func main() {
	r := chi.NewRouter()
	r.Get("/home", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFoundHandler()
	http.ListenAndServe(":3000", r)
}
