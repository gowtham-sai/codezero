package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSpec(t *testing.T) {
	t.Run("when no parsers registered", func(t *testing.T) {
		t.Run("should be able to register a registry", func(t *testing.T) {
			mockParser := &MockParser{}
			err := registerParser(mockParserType, mockParser)
			assert.NoError(t, err, "should not throw error")
		})
	})

	t.Run("when registry is already registered", func(t *testing.T) {
		mockParser := &MockParser{}
		t.Run("should not be able to register a registry", func(t *testing.T) {
			if _, found := p.registeredParsers[mockParserType]; !found {
				p.registeredParsers[mockParserType] = mockParser
			}
			err := registerParser(mockParserType, mockParser)
			assert.Errorf(t, err, "should throw error")

		})
	})
}
