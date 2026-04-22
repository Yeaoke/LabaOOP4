package http

import (
	"LabaOOP4/go-server/config"
	"bytes"
	"io"
	"net/http"
)

func Proxy(w http.ResponseWriter, r *http.Request, path string, method string) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read error", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(method, "http://localhost:8081"+path, bytes.NewReader(body))
	if err != nil {
		http.Error(w, "Create request error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Proxy error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func CallPython(r *http.Request, path string) ([]byte, int, error) {
	var bodyReader io.Reader
	if r.Body != nil && r.ContentLength > 0 {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, 0, err
		}
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(r.Method, "http://localhost:8081"+path, bodyReader)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, values := range r.Header {
		for _, value := range values {
			if key != "Host" && key != "Connection" {
				req.Header.Add(key, value)
			}
		}
	}

	resp, err := config.PythonClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, nil
}
