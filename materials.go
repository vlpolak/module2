package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("materials.html"))
var validPath = regexp.MustCompile("^/(/|/materials)/$")

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func materialsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "materials")
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.Execute(w, tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", makeHandler(rootHandler))
	http.HandleFunc("/materials/", makeHandler(materialsHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
