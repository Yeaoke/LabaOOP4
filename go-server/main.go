package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	backendURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal("Ошибка парсинга бэкенда: ", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(backendURL)

	r.PathPrefix("/api/").Handler(proxy)

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(fs)

	log.Fatal("localhost: ", "http://localhost:8080")
	log.Fatal("API: ", backendURL)
	log.Fatal(http.ListenAndServe(":3030", r))
}
