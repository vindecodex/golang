package controllers

import (
	"net/http"
	"text/template"
)

// Welcome
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseGlob("./views/*/*.html"))
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
