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
	r := mux.NewRouter()

	// Рендеринг страниц
	r.HandleFunc("/", html.PageHome)
	r.HandleFunc("/add", html.AddPageRendering)
	r.HandleFunc("/info/{id}", html.InfoPageRendering)
	r.HandleFunc("/edit/{id}", html.EditPageRendering)
	r.HandleFunc("/delete/{id}", html.DeletePageRendering)

	// Проксирование API
	r.HandleFunc("/api/add", handlers.AddHandler(pythonBackend)).Methods(http.MethodPost)
	r.HandleFunc("/api/edit/{id}", handlers.EditHandler(pythonBackend)).Methods(http.MethodPost)
	r.HandleFunc("/api/info/{id}", handlers.InfoHandler(pythonBackend)).Methods(http.MethodGet)
	r.HandleFunc("/api/delete/{id}", handlers.DeleteHandler(pythonBackend)).Methods(http.MethodDelete)

	proxy, err := handlers.NewProxy(pythonBackend)
	if err != nil {
		log.Fatal("Failed create proxy", err)
	}
	r.PathPrefix("/api/").Handler(http.StripPrefix("/api/", proxy))

	println("BFF (Go - Server): ", golangBFF)
	println("Backend: ", pythonBackend)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
