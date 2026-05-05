package htmlrendering

import (
	"LabaOOP4/go-server/cache"
	dto "LabaOOP4/go-server/dto"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Page struct {
	Title     string
	Companies []dto.IndustrialCompanies
	Company   *dto.IndustrialCompanies
}

var templates = template.Must(template.ParseFiles(
	"html/static/index.html",
	"html/static/add.html",
	"html/static/edit.html",
	"html/static/info.html",
	"html/static/delete.html",
))

func AddPageRendering(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InfoPageRendering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.NotFound(w, r)
	}

	company, ok := cache.CacheGet(id)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusNotFound)
	}

	data := Page{
		Title:   "Информация",
		Company: &company,
	}

	if err := templates.ExecuteTemplate(w, "info.html", data); err != nil {
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
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	company, ok := cache.CacheGet(id)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusNotFound)
	}

	data := Page{
		Title:   "Редактирование",
		Company: &company,
	}

	if err := templates.ExecuteTemplate(w, "edit.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeletePageRendering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	company, ok := cache.CacheGet(id)
	if !ok {
		http.Error(w, "Invalid ID", http.StatusNotFound)
		return
	}

	data := Page{
		Title:   "Удаление",
		Company: &company,
	}

	if err := templates.ExecuteTemplate(w, "delete.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
