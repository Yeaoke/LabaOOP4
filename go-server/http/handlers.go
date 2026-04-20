package http

import (
	"log"
	"net/http"
	"text/template"
)

type HTTPHandlers struct {
	httpHandlers *HTTPHandlers
}

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseGlob("template/index.html")
	if err != nil {
		log.Fatal(ErrorLaunchServer)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func proxyHandler(w http.ResponseWriter, r *http.Request) {

}
