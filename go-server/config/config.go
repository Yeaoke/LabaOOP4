package config

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

var PythonClient *http.Client

func InitHTTPClient() {
	PythonClient = &http.Client{
		Timeout: 30 * time.Second,

		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}
