package validation

import (
	cfg "LabaOOP4/go-server/config"
	"LabaOOP4/go-server/dto/models"
	"encoding/json"
	"net/http"
)

func ValidatorHandler(w http.ResponseWriter, r *http.Request) {
	var req models.IndustrialCompanies

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfg.Validate.Struct(req); err != nil {
		http.Error(w, "Invalid JSON"+err.Error(), http.StatusInternalServerError)
		return
	}
}
