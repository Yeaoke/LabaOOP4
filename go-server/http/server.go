package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
	server   *HTTPServer
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: handlers,
	}
}

func (s *HTTPServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	server := http.Server{
		Addr:    "8081",
		Handler: router,
	}

	defer server.Close()

	err := server.ListenAndServe()
	if errors.Is(err, ErrorLaunchServer) {
		return nil
	}

	return err
}
