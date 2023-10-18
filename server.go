package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.ListenAndServe(":4242", nil)
}
