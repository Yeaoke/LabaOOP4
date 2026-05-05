package http

import (
	"LabaOOP4/go-server/cache"
	cfg "LabaOOP4/go-server/config"
	dto "LabaOOP4/go-server/dto"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.IndustrialCompanies

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := cfg.Validate.Struct(&req); err != nil {
			http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
			return
		}

		if req.ID == uuid.Nil {
			req.ID = uuid.New()
		}

		jsonData, _ := json.Marshal(req)

		resp, err := http.Post(target+"/add", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		var response dto.IndustrialCompanies
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.WriteHeader(resp.StatusCode)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Failed to encode response: %v", err)
		}
	}
}

func InfoHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		// Проверяем кэш
		if company, ok := cache.CacheGet(id); ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(company)
			return
		}

		// Запрос к бэкенду (GET, без тела)
		targetURL := target + "/info/" + idStr
		resp, err := http.Get(targetURL)
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			return
		}

		var response dto.IndustrialCompanies
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func EditHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var req dto.IndustrialCompanies
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := cfg.Validate.Struct(&req); err != nil {
			http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
			return
		}

		req.ID = id // Принудительно ставим ID из URL

		jsonData, _ := json.Marshal(req)

		resp, err := http.Post(target+"/edit/"+idStr, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			return
		}

		var response dto.IndustrialCompanies
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		cache.CacheRemove(id)

		targetURL := target + "/delete/" + idStr
		reqBackend, err := http.NewRequest(http.MethodDelete, targetURL, nil)
		if err != nil {
			http.Error(w, "Failed to create backend request", http.StatusInternalServerError)
			return
		}
		reqBackend.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqBackend)
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(map[string]string{"message": "Deleted", "id": idStr})
	}
}
