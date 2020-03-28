package http

import (
	"log"

	"gopkg.in/yaml.v2"

	"codezero/deps"
)

func parseDependency(spec deps.Spec) (d Dependencies, err error) {
	err = yaml.Unmarshal([]byte(spec), &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	return
}


func parseScenario(spec deps.Spec) (s Scenario, err error) {
	err = yaml.Unmarshal([]byte(spec), &s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	return
}
