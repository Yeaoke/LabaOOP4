package main

import (
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	pythonBackend := "http://localhost:8080"

	r.HandleFunc("api/py", proxy.proxyHandler(pythonBackend))

}
