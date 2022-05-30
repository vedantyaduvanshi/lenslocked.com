package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func main() {

	homeView = views.NewView("khudse", "views/home.html")
	contactView = views.NewView("khudse", "views/contact.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	// Assets
	assetHandler := http.FileServer(http.Dir("./static/"))
	assetHandler = http.StripPrefix("/static/", assetHandler)
	r.PathPrefix("/static/").Handler(assetHandler)

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
