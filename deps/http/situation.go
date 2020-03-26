package http

import (
	"fmt"
	"net/http"
)

type (
	SituationName string
	Situations    map[SituationName]Situation

	Situation struct {
		Req Request  `yaml:"req"`
		Res Response `yaml:"res"`
	}

	Spec struct {
		Port int `yaml:"port"`
	}
)


func (s *Spec) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}


func (s *Situation) StartSituation(spec Spec) {
	handler := s.Res.createHandler()
	http.HandleFunc(s.Req.Path, handler)
	http.ListenAndServe(spec.Addr(), nil)
}
