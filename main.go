package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/home", homeHandler)
	r.NotFoundHandler()
	http.ListenAndServe(":3000", r)
}
