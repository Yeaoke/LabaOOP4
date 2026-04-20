package main

import (
	"LabaOOP4/go-server/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	pythonBackend := "http://localhost:8080"

	r.HandleFunc("api/py", http.ProxyHandler(pythonBackend))

}
