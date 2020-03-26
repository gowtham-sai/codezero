package http

import (
	"codezero/deps"
)

const (
	Type deps.Type = "http"

	Get     Method = "GET"
	Head    Method = "HEAD"
	Post    Method = "POST"
	Put     Method = "PUT"
	Patch   Method = "PATCH" // RFC 5789
	Delete  Method = "DELETE"
	Connect Method = "CONNECT"
	Options Method = "OPTIONS"
	Trace   Method = "TRACE"
)

var (
	registeredDependencies = Handler{}
)

func main() {}

type (
	Handler   map[ServiceName]Dependency
	Situations map[SituationName]Situation

	ServiceName   string
	SituationName string

	Method string
	Header map[string][]string
	Query  map[string]string

	Request struct {
		Method  Method `yaml:"method"`
		Path    string `yaml:"path"`
		Query   Query  `yaml:"query"`
		Headers Header `yaml:"headers"`
	}

	Response struct {
		Body    string `yaml:"body"`
		Headers Header `yaml:"headers"`
	}

	Situation struct {
		Req Request  `yaml:"req"`
		Res Response `yaml:"res"`
	}

	Dependency struct {
		Sits Situations `yaml:"situations"`
	}

	Spec struct {
		Port int `yaml:"port"`
	}
)

func (s *Situation) StartSituation(spec Spec) {
}

