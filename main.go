package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Welcome to awesome Site </h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" {
		fmt.Fprint(w, "To get in touch , Send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> Cant find page fk u </h1>")
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", handlerFunc)
	http.ListenAndServe(":3000", r)
}
