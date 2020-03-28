package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSpecAddrShouldReturnPort(t *testing.T) {
	s := Spec{Port: 8010}
	assert.Equal(t, ":8010", s.Addr())
}

func TestStopSituation(t *testing.T) {
	s := &Situation{}
	assert.NotPanics(t, func(){s.StopSituation()})
}

func TestShouldStartSituation(t *testing.T) {
	situation := Situation{
		Req: &Request{
			Method:  Get,
			Path:    "/v1/ping",
			Query:   Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
			Headers: Header{"Accept-Encoding": []string{"gzip", "compress"}},
		},
		Res: &Response{
			Code:    http.StatusAccepted,
			Body:    fmt.Sprintf("%s\n", `{"ping": "pong"}`),
			Headers: Header{"Accept-Encoding": []string{"application/json", "gzip"}},
		},
	}
	spec := Spec{Port: 8010}
	situation.StartSituation(spec)
	defer situation.StopSituation()

	req, err := http.NewRequest(string(Get), fmt.Sprintf("http://localhost:%d/v1/ping", spec.Port), nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
	assert.Equal(t, situation.Res.Headers["Accept-Encoding"], resp.Header["Accept-Encoding"])
	assert.Equal(t, string(respBytes), fmt.Sprintf("%s\n", `{"ping": "pong"}`))
}
