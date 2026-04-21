package http

import (
	cfg "LabaOOP4/go-server/config"
	"LabaOOP4/go-server/models"
	"encoding/json"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req models.IndustrialCompanies

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := cfg.Validate.Struct(&req); err != nil {
		http.Error(w, "Validation error"+err.Error(), http.StatusBadRequest)
		return
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {

}

func EditHandler(w http.ResponseWriter, r *http.Request) {

}
