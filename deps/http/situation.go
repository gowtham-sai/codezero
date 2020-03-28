package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type (
	SituationName string
	Situations    map[SituationName]*Situation

	Situation struct {
		Req *Request  `yaml:"req"`
		Res *Response `yaml:"res"`
		srv *http.Server
	}

	Spec struct {
		Port int `yaml:"port"`
	}
)

func (s *Spec) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Situation) StartSituation(spec Spec) {
	mux := http.NewServeMux()
	handler := s.Res.createHandler()
	mux.HandleFunc(s.Req.Path, handler)
	s.srv = &http.Server{
		Addr:    spec.Addr(),
		Handler: handler,
	}
	go func() {
		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Failed to start situation: %+v", err)
		}
	}()
}

func (s *Situation) StopSituation() (err error) {
	if s.srv != nil {
		s.srv.Shutdown(context.Background())
	}
	return
}
