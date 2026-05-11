package htmlrendering

import (
	"LabaOOP4/go-server/cache"
	dto "LabaOOP4/go-server/dto"
	"encoding/json"
	"html/template"
	"log"
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
	"html/static/holding.html",
))

// backendURL задаётся из main через Init
var backendURL string

func Init(target string) {
	backendURL = target
}

func fetchFromBackend(idStr string) (dto.IndustrialCompanies, bool) {
	resp, err := http.Get(backendURL + "/info/" + idStr)
	if err != nil || resp.StatusCode != http.StatusOK {
		return dto.IndustrialCompanies{}, false
	}
	defer resp.Body.Close()

	var company dto.IndustrialCompanies
	if err := json.NewDecoder(resp.Body).Decode(&company); err != nil {
		return dto.IndustrialCompanies{}, false
	}
	return company, true
}

func getCompany(id uuid.UUID, idStr string) (dto.IndustrialCompanies, bool) {
	if company, ok := cache.CacheGet(id); ok {
		return company, true
	}
	company, ok := fetchFromBackend(idStr)
	if ok {
		cache.CacheAdd(id, company)
	}
	return company, ok
}

func render(w http.ResponseWriter, tmpl string, data any) {
	if err := templates.ExecuteTemplate(w, tmpl, data); err != nil {
		log.Printf("[render] %s: %v", tmpl, err)
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func PageHome(backendTarget string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		resp, err := http.Get(backendTarget + "/")
		if err != nil {
			render(w, "index.html", Page{Title: "Главная"})
			return
		}
		defer resp.Body.Close()

		var companies []dto.IndustrialCompanies
		if err := json.NewDecoder(resp.Body).Decode(&companies); err != nil {
			render(w, "index.html", Page{Title: "Главная"})
			return
		}

		for _, c := range companies {
			if c.ID != uuid.Nil {
				cache.CacheAdd(c.ID, c)
			}
		}

		render(w, "index.html", Page{Title: "Главная", Companies: companies})
	}
}

func AddPageRendering(w http.ResponseWriter, r *http.Request) {
	render(w, "add.html", nil)
}

func HoldingPageRendering(w http.ResponseWriter, r *http.Request) {
	holdingName := r.URL.Query().Get("holding_name")

	if holdingName == "" {
		http.Error(w, "Required name for holding", http.StatusBadRequest)
		return
	}

	apiURL := backendURL + "/holding?holding_name=" + holdingName

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[HoldingPage] Error fetching from backend: %v", err)
		render(w, "holding.html", Page{Title: holdingName, Companies: []dto.IndustrialCompanies{}})
		return
	}
	defer resp.Body.Close() 

	var companies []dto.IndustrialCompanies

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&companies); err != nil {
			log.Printf("[HoldingPage] JSON Decode failed: %v", err)
			companies = []dto.IndustrialCompanies{} 

			for _, c := range companies {
				if c.ID != uuid.Nil {
					cache.CacheAdd(c.ID, c)
				}
			}
		}
	} else {
		
		log.Printf("[HoldingPage] Backend returned status: %d", resp.StatusCode)
		companies = []dto.IndustrialCompanies{}
	}

	render(w, "holding.html", Page{
		Title:     holdingName,
		Companies: companies,
	})
}

func InfoPageRendering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	company, ok := getCompany(id, idStr)
	if !ok {
		http.Error(w, "Компания не найдена", http.StatusNotFound)
		return
	}

	render(w, "info.html", Page{Title: "Информация", Company: &company})
}

func EditPageRendering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	company, ok := getCompany(id, idStr)
	if !ok {
		http.Error(w, "Компания не найдена", http.StatusNotFound)
		return
	}

	render(w, "edit.html", Page{Title: "Редактирование", Company: &company})
}

func DeletePageRendering(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	company, ok := getCompany(id, idStr)
	if !ok {
		http.Error(w, "Компания не найдена", http.StatusNotFound)
		return
	}

	render(w, "delete.html", Page{Title: "Удаление", Company: &company})
}