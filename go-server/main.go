package main

import (
	cfg "LabaOOP4/go-server/config"
	handlers "LabaOOP4/go-server/http"
	"LabaOOP4/go-server/validation"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	cfg.Validate = validator.New()

	r := mux.NewRouter()

	pythonBackend := "http://localhost:8081"

	r.HandleFunc("/add", validation.ValidatorHandler)
	r.HandleFunc("/", handlers.ProxyHandler(pythonBackend))

	println("API Gateway :8080")
	println("Backend: ", pythonBackend)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
