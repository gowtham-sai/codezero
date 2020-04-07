package deps

import "github.com/stretchr/testify/mock"

// MockHandler provides mock for dependency Handler interface
type MockHandler struct {
	mock.Mock
}

// StartSituation  mocks StartSituation of dependency Handler interface
func (m *MockHandler) StartSituation(spec Spec) error {
	args := m.Called(spec)
	return args.Error(0)
}

// StopSituation mocks StopSituation of dependency Handler interface
func (m *MockHandler) StopSituation(spec Spec) error {
	args := m.Called(spec)
	return args.Error(0)
}

// ParseSpec mocks ParseSpec of dependency Handler interface
func (m *MockHandler) ParseSpec(spec Spec) error {
	args := m.Called(spec)
	return args.Error(0)
}
