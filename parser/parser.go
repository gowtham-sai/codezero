package parser

import (
	"errors"
	"sync"

	"codezero/deps"
)

type registry map[parserType]Parser

var (
	p *parserRegistry

	parserAlreadyRegister = errors.New("parser_already_registered")
)

type spec string
type parserType string

type Parser interface {
	ParseSpec(spec spec) (dep deps.Dependency, err error)
}

type parserRegistry struct {
	lock              sync.Mutex
	registeredParsers registry
}

func registerParser(parserType parserType, parser Parser) (err error) {
	for pt := range p.registeredParsers {
		if pt == parserType {
			return parserAlreadyRegister
		}
	}
	p.registeredParsers[parserType] = parser
	return
}

func (p *registry) ParseSpec(spec spec) (dep deps.Dependency, err error) {
	return
}

func NewParser() (parserRegistry *parserRegistry) {
	return p
}

func init() {
	p = &parserRegistry{
		registeredParsers: make(registry),
	}
}
