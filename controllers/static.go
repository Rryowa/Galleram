package controllers

import (
	"net/http"
)

// StaticTemplate - Closure, for accessing data that typically isn’t available
// without using a global variable.
/*
Now we can write handler functions as if they had access to a Template object
while still returning a function with the signature that http.HandleFunc()
expects. This allows us to bypass the fact that http.HandleFunc()
doesnt permit us passing in custom variables without resorting
to global variables.
*/
func StaticTemplate(t Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
