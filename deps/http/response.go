package http

import (
	"net/http"
)

type Response struct {
	StatusCode int    `yaml:"status_code"`
	Body       string `yaml:"body"`
	Headers    Header `yaml:"headers"`
}

func (resp *Response) createHandler() http.HandlerFunc {
	responseWriter := func(w http.ResponseWriter, r *http.Request) {
		for hKey, hValues := range resp.Headers {
			w.Header()[hKey] = hValues
		}
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(resp.Body))
	}
	return responseWriter
}
