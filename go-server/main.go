package main

import (
	cfg "LabaOOP4/go-server/config"
	html "LabaOOP4/go-server/html/html-rendering"
	handlers "LabaOOP4/go-server/http"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	cfg.Validate = validator.New()

	r := mux.NewRouter()

	// Часть прокси
	pythonBackend := "http://localhost:8081"

	target, err := url.Parse(pythonBackend)
	if err != nil {

	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	r.PathPrefix("/api/py").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("BFF: %s %s -> Proxy to Python", r.Method, r.URL.Path)

		r.Host = target.Host
		proxy.ServeHTTP(w, r)
	}))
	//

	r.HandleFunc("/", html.PageHome)
	r.HandleFunc("/add", handlers.AddHandler)
	r.HandleFunc("/info", handlers.InfoHandler)
	r.HandleFunc("/edit", handlers.EditHandler)

	println("API Gateway :8080")
	println("Backend: ", pythonBackend)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
