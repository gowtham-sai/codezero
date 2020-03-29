package http

import (
	"gopkg.in/yaml.v2"

	"codezero/deps"
)

func parseDependency(spec deps.Spec) (d dependencies, err error) {
	err = yaml.Unmarshal([]byte(spec), &d)
	return
}

func parseScenario(spec deps.Spec) (s scenario, err error) {
	err = yaml.Unmarshal([]byte(spec), &s)
	return
}
