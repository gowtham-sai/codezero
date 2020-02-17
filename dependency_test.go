package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDependency(t *testing.T) {
	file, err := ioutil.TempFile("./", ".codezero_test.*.yaml")
	require.NoError(t, err)
	defer os.Remove(file.Name())
	testData :=
		`
req:
  method: GET
  path: /transport/v1/estimate
  query:
    waypoints: 102.6123,-6.1235|102.113,-6.2343
  headers:
    Accept-Encoding:
      - gzip
      - compress
res:
  body: |
    {"ping": "pong"}
`

	file.Write([]byte(testData))

	dep, err := ParseDependency(file.Name())
	assert.NoError(t, err)

	assert.Equal(t, "GET", string(dep.Req.Method))
	assert.Equal(t, "/transport/v1/estimate", string(dep.Req.Path))
	assert.Equal(t, Query{"waypoints": "102.6123,-6.1235|102.113,-6.2343"}, dep.Req.Query)
	assert.Equal(t, Header{"Accept-Encoding": []string{"gzip", "compress"}}, dep.Req.Headers)

	assert.Equal(t, `{"ping": "pong"}`, strings.TrimSpace(dep.Res.Body))
}
