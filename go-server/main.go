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

	proxy, err := handlers.NewProxy(pythonBackend)
	if err != nil {
		log.Fatal("Failed create proxy", err)
	}
	r := mux.NewRouter()

	// Рендеринг static
	r.HandleFunc("/", html.PageHome)
	r.HandleFunc("/add", html.AddPageRendering)
	r.HandleFunc("/info", html.InfoPageRendering)
	r.HandleFunc("/edit", html.EditPageRendering)

	// Проксирование handlers
	r.PathPrefix("/api/").Handler(http.StripPrefix("/api/", proxy))

	println("BFF (Go - Server): 8080")
	println("Backend: ", pythonBackend)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
