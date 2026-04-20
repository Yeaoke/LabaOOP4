package http

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"text/template"
)

type HTTPHandlers struct {
	httpHandlers *HTTPHandlers
}

var templates = template.Must(template.ParseGlob("static/*html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add" {
		http.NotFound(w, r)
		return
	}

	err := templates.ExecuteTemplate(w, "add.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info" {
		http.NotFound(w, r)
		return
	}

	err := templates.ExecuteTemplate(w, "info.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	err := templates.ExecuteTemplate(w, "edit.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func proxyHandler(targetURL string) http.HandlerFunc {
	url, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/py")

		r.Host = url.Host

		log.Printf("Proxying request: %s %s -> %s%s", r.Method, r.URL.Path, url.Host, r.URL.Path)

		proxy.ServeHTTP(w, r)
	}
}
