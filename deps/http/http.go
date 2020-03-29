package http

import (
	"codezero/deps"
)

const (
	depType deps.Type = "http"

	methodGet     method = "GET"
	methodHead    method = "HEAD"
	methodPost    method = "POST"
	methodPut     method = "PUT"
	methodPatch   method = "PATCH" // RFC 5789
	methodDelete  method = "DELETE"
	methodConnect method = "CONNECT"
	methodOptions method = "OPTIONS"
	methodTrace   method = "TRACE"
)

var (
	registeredDependencies = &handler{}
)

type (
	serviceName string

	method string
	header map[string][]string
	query  map[string]string
)

func init() {
	deps.RegisterHandler(depType, registeredDependencies)
}
