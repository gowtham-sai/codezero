package http

import (
	"log"

	"gopkg.in/yaml.v2"

	"codezero/deps"
)

func parseSpec(spec deps.Spec) (d Dependencies, err error) {
	err = yaml.Unmarshal([]byte(spec), &d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return
	}
	return
}
