package deps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockHandler(t *testing.T) {
	t.Run("mock handler start situation", func(t *testing.T) {
		t.Run("should return error", func(t *testing.T) {
			mockHandler := MockHandler{}
			mockHandler.On("StartSituation", Spec(``)).Return(fmt.Errorf("situation error"))
			assert.Error(t, mockHandler.StartSituation(``))
		})

		t.Run("should return error", func(t *testing.T) {
			mockHandler := MockHandler{}
			mockHandler.On("StartSituation", Spec(``)).Return(nil)
			assert.NoError(t, mockHandler.StartSituation(``))
		})
	})
}
