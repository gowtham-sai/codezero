package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"codezero/deps"
)

func TestHandler(t *testing.T) {
	testSVC := serviceName("service_xyz")
	testSit2xx := situationName("response_2xx")
	serviceXYZSpec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz.yml")
	require.NoError(t, err, "yamlFile.Get error")

	t.Run("parse spec", func(t *testing.T) {
		t.Run("if service name already present", func(t *testing.T) {
			h := &handler{Deps: dependencies{testSVC: &dependency{Sits: situations{testSit2xx: &situation{}}}}}
			t.Run("should append situations", func(t *testing.T) {
				spec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz_response_5xx.yml")
				require.NoError(t, err, "yamlFile.Get error")

				h.ParseSpec(deps.Spec(spec))
				assert.Equal(t, 2, len(h.Deps[testSVC].Sits))
			})
		})

		t.Run("if service name not present", func(t *testing.T) {
			t.Run("should store situations", func(t *testing.T) {
				h := &handler{Deps: dependencies{}}
				h.ParseSpec(deps.Spec(serviceXYZSpec))

				assert.Equal(t, &dependency{
					Sits: situations{
						testSit2xx: &situation{
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
				}, h.Deps[testSVC])
			})
		})
	})

	t.Run("stop situation", func(t *testing.T) {
		h := &handler{Deps: dependencies{}}
		h.ParseSpec(deps.Spec(serviceXYZSpec))
		t.Run("situation should not be reachable", func(t *testing.T) {
			startSituationErr := h.StartSituation(`
service_xyz:
  response_2xx:
    port: 8010
`)
			assert.NoError(t, startSituationErr)
			req, err := http.NewRequest(string(methodGet), fmt.Sprintf("http://localhost:%d/v1/ping", 8010), nil)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			assert.Equal(t, h.Deps[testSVC].Sits[testSit2xx].Res.StatusCode(), resp.StatusCode)

			h.StopSituation(`
service_xyz:
  response_2xx:
    port: 8010
`)
			retriedRes, err := http.DefaultClient.Do(req)
			require.Error(t, err)
			assert.Nil(t, retriedRes)
		})
	})

	t.Run("start situation", func(t *testing.T) {
		h := &handler{Deps: dependencies{}}
		h.ParseSpec(deps.Spec(serviceXYZSpec))
		t.Run("should start a situation given a spec", func(t *testing.T) {
			defer h.StopSituation(`
service_xyz:
  response_2xx:
    port: 8010
`)
			startSituationErr := h.StartSituation(`
service_xyz:
  response_2xx:
    port: 8010
`)
			assert.NoError(t, startSituationErr)
			req, err := http.NewRequest(string(methodGet), fmt.Sprintf("http://localhost:%d/v1/ping", 8010), nil)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, h.Deps[testSVC].Sits[testSit2xx].Res.Headers["Accept-Encoding"], resp.Header["Accept-Encoding"])
			assert.Equal(t, string(respBytes), fmt.Sprintf("%s\n", `{"ping": "pong"}`))
		})
	})
}
