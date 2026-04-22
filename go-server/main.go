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

	cfg.InitHTTPClient()

	r := mux.NewRouter()

	r.HandleFunc("/", html.PageHome)
	r.HandleFunc("/add", handlers.AddHandler)
	r.HandleFunc("/info", handlers.InfoHandler)
	r.HandleFunc("/edit", handlers.EditHandler)

	println("API Gateway :8080")
	println("Backend: ", "http://localhost:8081")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
