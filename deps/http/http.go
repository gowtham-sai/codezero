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
	registeredDependencies = &handler{}
)

type (
	ServiceName string

	Method string
	Header map[string][]string
	Query  map[string]string
)

func init() {
	deps.RegisterHandler(Type, registeredDependencies)
}
