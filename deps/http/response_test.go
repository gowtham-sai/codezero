package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	resp := &Response{
		StatusCode: http.StatusAccepted,
		Body:       `{"ping": "pong"}`,
		Headers:    Header{"Accept-Encoding": []string{"application/json", "gzip"}},
	}
	handlerFunc := resp.createHandler()

	r := httptest.NewRequest(http.MethodGet, "/v1/ping", nil)
	w := httptest.NewRecorder()

	handlerFunc(w, r)

	assert.Equal(t, resp.StatusCode, w.Code)
	assert.Equal(t, []byte(resp.Body), w.Body.Bytes())
	assert.Equal(t, resp.Headers["Accept-Encoding"], w.Header()["Accept-Encoding"])
}
