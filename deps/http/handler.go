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

func (h *Handler) ParseSpec(spec deps.Spec) error {
	panic("implement me")
}
