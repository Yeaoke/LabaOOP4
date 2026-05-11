package http

import (
	"LabaOOP4/go-server/cache"
	cfg "LabaOOP4/go-server/config"
	dto "LabaOOP4/go-server/dto"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func FilterHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Parameter 'q' is required", http.StatusBadRequest)
			return
		}

		resp, err := http.Get(target + "/")
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		var allCompanies []dto.IndustrialCompanies
		if err := json.NewDecoder(resp.Body).Decode(&allCompanies); err != nil {
			http.Error(w, "Failed to decode companies", http.StatusInternalServerError)
			return
		}

		query = strings.ToLower(query)
		var filtered []dto.IndustrialCompanies

		for _, c := range allCompanies {
			if strings.Contains(strings.ToLower(c.CompanyName), query) {
				filtered = append(filtered, c)
				if c.ID != uuid.Nil {
					cache.CacheAdd(c.ID, c)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filtered)
	}
}

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

		jsonData, err := json.Marshal(req)
		if err != nil {
			http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(target+"/add", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			log.Printf("[AddHandler] backend error %d: %s", resp.StatusCode, string(body))
			http.Error(w, "Backend error: "+string(body), resp.StatusCode)
			return
		}

		var response dto.IndustrialCompanies
		if err := json.Unmarshal(body, &response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
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

		if company, ok := cache.CacheGet(id); ok {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(company)
			return
		}

		resp, err := http.Get(target + "/info/" + idStr)
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			w.Write(body)
			return
		}

		var response dto.IndustrialCompanies
		if err := json.Unmarshal(body, &response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
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

		req.ID = id

		jsonData, err := json.Marshal(req)
		if err != nil {
			http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post(target+"/edit/"+idStr, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			log.Printf("[EditHandler] backend error %d: %s", resp.StatusCode, string(body))
			http.Error(w, "Backend error: "+string(body), resp.StatusCode)
			return
		}

		var response dto.IndustrialCompanies
		if err := json.Unmarshal(body, &response); err != nil {
			http.Error(w, "Failed to decode response", http.StatusInternalServerError)
			return
		}

		cache.CacheAdd(response.ID, response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
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

		reqBackend, err := http.NewRequest(http.MethodDelete, target+"/delete/"+idStr, nil)
		if err != nil {
			http.Error(w, "Failed to create backend request", http.StatusInternalServerError)
			return
		}
		reqBackend.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(reqBackend)
		if err != nil {
			http.Error(w, "Failed to reach backend", http.StatusBadGateway)
			return
		}
		defer response.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(map[string]string{"message": "Deleted", "id": idStr})
	}
}