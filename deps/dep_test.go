package deps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	t.Run("register new type Handler", func(t *testing.T) {
		t.Run("should be able to register", func(t *testing.T) {
			mockHandler := &MockHandler{}
			assert.NoError(t, RegisterHandler(Type("dummy"), mockHandler))
		})
	})
}
