package http

import (
	"net/http"
)

type response struct {
	Code    int    `yaml:"status_code"`
	Body    string `yaml:"body"`
	Headers header `yaml:"headers"`
}

func (resp *response) createHandler() http.HandlerFunc {
	responseWriter := func(w http.ResponseWriter, r *http.Request) {
		for hKey, hValues := range resp.Headers {
			w.Header()[hKey] = hValues
		}
		w.WriteHeader(resp.StatusCode())
		w.Write([]byte(resp.Body))
	}
	return responseWriter
}

func (resp *response) StatusCode() int {
	if resp.Code != 0 {
		return resp.Code
	}
	return http.StatusOK
}
