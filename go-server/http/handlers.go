package http

import (
	"LabaOOP4/go-server/cache"
	cfg "LabaOOP4/go-server/config"
	dto "LabaOOP4/go-server/dto"
	"bytes"
	"log"

	//"LabaOOP4/go-server/dto"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func AddHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.IndustrialCompanies

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

		jsonData, _ := json.Marshal(req)

		resp, err := http.Post(target, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed read buffer", http.StatusBadGateway)
		}

		defer resp.Body.Close()

		var response dto.IndustrialCompanies

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed decode", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Failed to encode response: %v", err)
			return
		}
	}
}

// Тут надо всё менять
func InfoHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.IndustrialCompanies

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
			return
		}

		if req.ID == uuid.Nil {
			req.ID = uuid.New()
		}

		cache.CacheRemove(req.ID)
		cache.CacheCheck(req.ID)

		targetURL := target + "/info"

		jsonData, _ := json.Marshal(req)

		resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed get answer from server", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		var response dto.IndustrialCompanies

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed decode", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Failed to encode response: %v", err)
			return
		}
	}
}

func EditHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.IndustrialCompanies

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

		cache.CacheRemove(req.ID)
		cache.CacheCheck(req.ID)

		targetURL := target + "/edit"

		jsonData, _ := json.Marshal(req)

		resp, err := http.Post(targetURL, "aplication/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed to get answer from server", http.StatusBadGateway)
			return
		}
		resp.Body.Close()

		var response dto.IndustrialCompanies

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed decode", http.StatusBadRequest)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Failed to encode response: %v", err)
			return
		}
	}
}

func DeleteHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.IndustrialCompanies

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON"+err.Error(), http.StatusBadRequest)
			return
		}

		if req.ID == uuid.Nil {
			req.ID = uuid.New()
		}

		cache.CacheRemove(req.ID)

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(req); err != nil {
			log.Printf("Failed to encode response: %v", err)
			return
		}
	}
}
