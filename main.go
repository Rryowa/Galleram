package main

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"path/filepath"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		http.Error(w, "wrong template path", http.
			StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "passed wrong data", http.
			StatusInternalServerError)
		return
	}

}

func main() {
	r := chi.NewRouter()
	r.Get("/home", homeHandler)

	r.NotFoundHandler()
	http.ListenAndServe(":3000", r)
}
