package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	resp := &Response{
		Code:    http.StatusAccepted,
		Body:    `{"ping": "pong"}`,
		Headers: Header{"Accept-Encoding": []string{"application/json", "gzip"}},
	}
	handlerFunc := resp.createHandler()

	r := httptest.NewRequest(http.MethodGet, "/v1/ping", nil)
	w := httptest.NewRecorder()

	handlerFunc(w, r)

	assert.Equal(t, resp.StatusCode(), w.Code)
	assert.Equal(t, []byte(resp.Body), w.Body.Bytes())
	assert.Equal(t, resp.Headers["Accept-Encoding"], w.Header()["Accept-Encoding"])
}

func TestStatusCode(t *testing.T) {
	t.Run("when status code is not given", func(t *testing.T) {
		resp := &Response{}
		t.Run("should return StatusOK as default code", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, resp.StatusCode())
		})
	})

	t.Run("when status code is given", func(t *testing.T) {
		resp := &Response{Code: http.StatusAccepted}
		t.Run("should return given status code", func(t *testing.T) {
			assert.Equal(t, http.StatusAccepted, resp.StatusCode())
		})
	})
}
