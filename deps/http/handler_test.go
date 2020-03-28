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
	testSVC := ServiceName("service_xyz")
	testSit2xx := SituationName("response_2xx")
	serviceXYZSpec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz.yml")
	require.NoError(t, err, "yamlFile.Get error")

	t.Run("parse spec", func(t *testing.T) {
		t.Run("if service name already present", func(t *testing.T) {
			h := &Handler{Deps: Dependencies{testSVC: &Dependency{Sits: Situations{testSit2xx: &Situation{}}}}}
			t.Run("should append situations", func(t *testing.T) {
				spec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz_response_5xx.yml")
				require.NoError(t, err, "yamlFile.Get error")

				h.ParseSpec(deps.Spec(spec))
				assert.Equal(t, 2, len(h.Deps[testSVC].Sits))
			})
		})

		t.Run("if service name not present", func(t *testing.T) {
			t.Run("should store situations", func(t *testing.T) {
				h := &Handler{Deps: Dependencies{}}
				h.ParseSpec(deps.Spec(serviceXYZSpec))

				assert.Equal(t, &Dependency{
					Sits: Situations{
						testSit2xx: &Situation{
							Req: &Request{
								Method:  Get,
								Path:    "/v1/ping",
								Query:   Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
								Headers: Header{"Accept-Encoding": []string{"gzip", "compress"}},
							},
							Res: &Response{
								Body: fmt.Sprintf("%s\n", `{"ping": "pong"}`),
							},
						},
					},
				}, h.Deps[testSVC])
			})
		})
	})
}
