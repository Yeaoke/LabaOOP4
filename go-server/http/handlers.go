package http

import (
	"LabaOOP4/go-server/cache"
	cfg "LabaOOP4/go-server/config"
	"log"

	//"LabaOOP4/go-server/dto"

	"LabaOOP4/go-server/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

	if req.ID == uuid.Nil {
		req.ID = uuid.New()
	}

	cache.Caching(req.ID, req)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		log.Printf("Failed to encode response: %v", err)
		return
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	//var req dto.IndustrialCompaniesResponse

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
