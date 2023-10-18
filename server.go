package main

import (
	"embed"
	"net/http"
	"html/template"
)

//go:embed index.html
var content embed.FS
func main() {

		tmpl,err := template.ParseFS(content, "index.html")
		if err != nil {
			panic(err)
		}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Failed to execute.", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":4242", nil)
}
