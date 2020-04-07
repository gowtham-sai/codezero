package deps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("registering new type Handler", func(t *testing.T) {
		t.Run("should be able to register", func(t *testing.T) {
			mockHandler := &MockHandler{}
			assert.NoError(t, RegisterHandler(Type("dummy"), mockHandler))
		})
	})

	t.Run("registering existing type Handler", func(t *testing.T) {
		handlers = map[Type]Handler{}
		t.Run("should not be able to register", func(t *testing.T) {
			mockHandler := &MockHandler{}
			assert.NoError(t, RegisterHandler(Type("mocked"), mockHandler))
			assert.EqualError(t, RegisterHandler(Type("mocked"), mockHandler), ErrTypeRegistered.Error())
		})
	})
}
