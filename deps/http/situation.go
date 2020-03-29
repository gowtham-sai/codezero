package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type (
	situationName string
	situations    map[situationName]*situation

	situation struct {
		Req *request  `yaml:"req"`
		Res *response `yaml:"res"`
		srv *http.Server
	}

	spec struct {
		Port int `yaml:"port"`
	}
)

func (s *spec) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *situation) StartSituation(spec spec) {
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

func (s *situation) StopSituation() (err error) {
	if s.srv != nil {
		s.srv.Shutdown(context.Background())
	}
	return
}
