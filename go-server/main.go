package main

import (
	cfg "LabaOOP4/go-server/config"
	html "LabaOOP4/go-server/html/html-rendering"
	handlers "LabaOOP4/go-server/http"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	cfg.Validate = validator.New()

	pythonBackend := "http://localhost:8081"
	golangBFF := "http://localhost:8080"

	html.Init(pythonBackend)

	r := mux.NewRouter()

	// Рендеринг HTML-страниц
	r.HandleFunc("/", html.PageHome(pythonBackend)).Methods(http.MethodGet)
	r.HandleFunc("/add", html.AddPageRendering).Methods(http.MethodGet)
	r.HandleFunc("/info/{id}", html.InfoPageRendering).Methods(http.MethodGet)
	r.HandleFunc("/edit/{id}", html.EditPageRendering).Methods(http.MethodGet)
	r.HandleFunc("/delete/{id}", html.DeletePageRendering).Methods(http.MethodGet)
	r.HandleFunc("/holding", html.HoldingPageRendering).Methods(http.MethodGet)
	
	// API (проксирование и обработка)
	r.HandleFunc("/api/filter", handlers.FilterHandler(pythonBackend)).Methods(http.MethodGet)
	r.HandleFunc("/api/add", handlers.AddHandler(pythonBackend)).Methods(http.MethodPost)
	r.HandleFunc("/api/edit/{id}", handlers.EditHandler(pythonBackend)).Methods(http.MethodPost)
	r.HandleFunc("/api/info/{id}", handlers.InfoHandler(pythonBackend)).Methods(http.MethodGet)
	r.HandleFunc("/api/delete/{id}", handlers.DeleteHandler(pythonBackend)).Methods(http.MethodDelete)

	proxy, err := handlers.NewProxy(pythonBackend)
	if err != nil {
		log.Fatal("Failed to create proxy: ", err)
	}

	r.PathPrefix("/api/action/").Handler(http.StripPrefix("/api", proxy))

	log.Printf("BFF  (Go):     %s", golangBFF)
	log.Printf("Backend (Py):  %s", pythonBackend)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}