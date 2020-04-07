package deps

import "errors"

var (
	// ErrTypeRegistered means that a handler for that particular type is already registered
	// We do not allow duplicate handlers for same type
	ErrTypeRegistered = errors.New("deps.type.registered")

	handlers = map[Type]Handler{}
)

type (
	// Type indicates the Handler type
	// a type for http Handler might be http and for tcp might be tcp
	// specs allows only registered handlers
	Type string

	// Spec as the name says, it indicates the spec for Handler
	// For a given scenario, all the specs will be passed to registered handler
	Spec string

	// Handler defines the interface for specific dependencies
	// any dependencies must implement Handler interface in order
	// to be a dependency for a scenario
	Handler interface {
		StartSituation(spec Spec) error
		StopSituation(spec Spec) error
		ParseSpec(spec Spec) error
	}
)

// RegisterHandler provides a way for dependency to register
// it requires type and handler
// if type is already registered, it throws error
func RegisterHandler(t Type, h Handler) (err error) {
	if _, found := handlers[t]; found {
		return ErrTypeRegistered
	}
	handlers[t] = h
	return
}
