package htmlrendering

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("html/static/*.html"))

func AddPageRendering(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InfoPageRendering(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "info.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PageHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EditPageRendering(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "edit.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
