package parser

import (
	"errors"
	"sync"

	"codezero/deps"
)

type registry map[Type]Parser

var (
	p *Registry

	parserAlreadyRegister = errors.New("parser_already_registered")
)

type (
	Spec string
	Type string

	Parser interface {
		ParseSpec(spec Spec) (dep deps.Dependency, err error)
	}

	Registry struct {
		lock    sync.Mutex
		parsers registry
	}
)

func Register(parserType Type, parser Parser) (err error) {
	for pt := range p.parsers {
		if pt == parserType {
			return parserAlreadyRegister
		}
	}
	p.parsers[parserType] = parser
	return
}

func (p *registry) ParseSpec(spec Spec) (dep deps.Dependency, err error) {
	return
}

func NewRegistry() (registry *Registry) {
	return p
}

func init() {
	p = &Registry{
		parsers: make(registry),
	}
}
