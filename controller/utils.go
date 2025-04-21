package controller

import (
	"io"
	"net/http"
)

func NewRequest(method string, w http.ResponseWriter, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, ASANA_API_URL+path, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return req, nil
}
