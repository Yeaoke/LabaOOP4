package http

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"text/template"
)

var templates = template.Must(template.ParseGlob("static/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "add.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "info.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}
	if err := templates.ExecuteTemplate(w, "edit.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProxyHandler(targetURL string) http.HandlerFunc {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Invalid backend URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	return func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.URL.Path
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/py")
		if r.URL.Path == "" {
			r.URL.Path = "/"
		}
		r.Host = parsedURL.Host

		log.Printf("[PROXY] %s %s → %s%s", r.Method, originalPath, parsedURL.Host, r.URL.Path)
		proxy.ServeHTTP(w, r)
	}
}
