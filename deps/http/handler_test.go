package http

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"codezero/deps"
)

func TestHandler(t *testing.T) {
	t.Run("parse spec", func(t *testing.T) {
		h := &Handler{}
		spec, err := ioutil.ReadFile("../../fixtures/deps/http/service_xyz.yml")
		require.NoError(t, err, "yamlFile.Get error")

		h.ParseSpec(deps.Spec(spec))
		assert.Equal(t, Dependency{
					Sits: Situations{
						"response_2xx": Situation{
							Req: Request{
								Method:  Get,
								Path:    "/v1/ping",
								Query:   Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"},
								Headers: Header{"Accept-Encoding": []string{"gzip", "compress"}},
							},
							Res: Response{
								Body: fmt.Sprintf("%s\n", `{"ping": "pong"}`),
							},
						},
					},
				}, h.Deps["voucher_service"])
	})
}
