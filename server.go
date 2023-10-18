package main

import (
	"embed"
	"net/http"
)

//go:embed index.html
var indexHTML string
func main() {
	tmpl = template.Must(template.New("index").Parse(indexHTML)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	http.ListenAndServe(":4242", nil)
}
