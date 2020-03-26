package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type (
	SituationName string
	Situations    map[SituationName]Situation

	Situation struct {
		Req Request  `yaml:"req"`
		Res Response `yaml:"res"`
		srv http.Server
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
	s.srv = http.Server{
		Addr: spec.Addr(),
	}
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start situation: %+v", err)
		}
	}()
}

func (s *Situation) StopSituation() {
	s.srv.Shutdown(context.Background())
}
