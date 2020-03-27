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

	Dependencies map[ServiceName]Dependency
)

func (h *Handler) StopSituation(spec deps.Spec) error {
	panic("implement me")
}

func (h *Handler) StartSituation(spec deps.Spec) error {
	panic("implement me")

}

func (h *Handler) ParseSpec(spec deps.Spec) (err error) {
	d, err := parseSpec(spec)
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
