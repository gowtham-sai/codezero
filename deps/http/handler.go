package http

import (
	"codezero/deps"
)

type (
	handler struct {
		Deps dependencies
	}

	dependency struct {
		Sits Situations `yaml:"situations"`
	}

	dependencies map[ServiceName]*dependency
	scenario     map[ServiceName]map[SituationName]Spec
)

func (h *handler) StopSituation(spec deps.Spec) (err error) {
	scenario, err := parseScenario(spec)
	if err != nil {
		return
	}

	for serviceName, situations := range scenario {
		for situationName := range situations {
			if dep, found := h.Deps[serviceName]; found {
				if sit, found := dep.Sits[situationName]; found {
					sit.StopSituation()
				}
			}
		}
	}

	return
}

func (h *handler) StartSituation(spec deps.Spec) (err error) {
	scenario, err := parseScenario(spec)
	if err != nil {
		return
	}
	for serviceName, situations := range scenario {
		for situationName, scenarioSpec := range situations {
			if dep, found := h.Deps[serviceName]; found {
				if sit, found := dep.Sits[situationName]; found {
					sit.StartSituation(scenarioSpec)
				}
			}
		}
	}
	return
}

func (h *handler) ParseSpec(spec deps.Spec) (err error) {
	d, err := parseDependency(spec)
	if err != nil {
		return err
	}
	for serviceName, dependency := range d {
		for situationName, situation := range dependency.Sits {
			if dep, found := h.Deps[serviceName]; found {
				dep.Sits[situationName] = situation
				h.Deps[serviceName] = dep
			} else {
				h.Deps[serviceName] = dependency
			}
		}
	}
	return
}
