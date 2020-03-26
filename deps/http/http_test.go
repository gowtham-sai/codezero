package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldStartSituation(t *testing.T) {
	situation := Situation{
		Req: Request{
			Method:  Get,
			Path:    "/v1/ping",
			Query:   Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
			Headers: Header{"Accept-Encoding": []string{"gzip", "compress"}},
		},
		Res: Response{
			Body: fmt.Sprintf("%s\n", `{"ping": "pong"}`),
		},
	}
	spec := Spec{Port: 8010}
	situation.StartSituation(spec)

	req, err := http.NewRequest(string(Get), fmt.Sprintf("http://localhost:%d/v1/ping", spec.Port), nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, string(respBytes), `{"ping": "pong"}`)
}
