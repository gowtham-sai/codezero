package http

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"codezero/deps"
)

func TestParseDependency(t *testing.T) {
	spec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz.yml")
	require.NoError(t, err, "yamlFile.Get error")

	actualHandler, err := parseDependency(deps.Spec(spec))
	require.NoError(t, err, "http.parseDependency error")

	expectedHandler := handler{
		Deps: dependencies{
			"service_xyz": &dependency{
				Sits: situations{
					"response_2xx": &situation{
						Req: &request{
							Method:  methodGet,
							Path:    "/v1/ping",
							Query:   query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
							Headers: header{"Accept-Encoding": []string{"gzip", "compress"}},
						},
						Res: &response{
							Body: fmt.Sprintf("%s\n", `{"ping": "pong"}`),
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expectedHandler.Deps, actualHandler)
}

func TestParseScenario(t *testing.T) {
	actualScenario, err := parseScenario(`
service_xyz:
  response_2xx:
    port: 8010
`)
	require.NoError(t, err)

	expectedScenario := scenario{"service_xyz": map[situationName]spec{"response_2xx": {Port: 8010}}}
	assert.Equal(t, expectedScenario, actualScenario)
}
