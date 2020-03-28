package http

import (
	"codezero/deps"
)

type (
	Handler struct {
		Deps Dependencies
	}

	Dependency struct {
		Sits Situations `yaml:"situations"`
	}

	Dependencies map[ServiceName]*Dependency
	Scenario     map[ServiceName]map[SituationName]Spec
)

func (h *Handler) StopSituation(spec deps.Spec) (err error) {
	scenario, err := parseScenario(spec)
	if err != nil {
		return
	}

	for serviceName, situations := range scenario {
		for situationName, _ := range situations {
			if dep, found := h.Deps[serviceName]; found {
				if sit, found := dep.Sits[situationName]; found {
					sit.StopSituation()
				}
			}
		}
	}

	return
}

func (h *Handler) StartSituation(spec deps.Spec) (err error) {
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

func (h *Handler) ParseSpec(spec deps.Spec) (err error) {
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
