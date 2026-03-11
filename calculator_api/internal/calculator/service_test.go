package calculator

import (
	"shared/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Stubs
type mockParser struct {
	result float64
	err    error
}

func (m *mockParser) Evaluate(expression string) (float64, error) {
	return m.result, m.err
}

type mockPublisher struct {
	published bool
	err       error
}

func (m *mockPublisher) Publish(message models.CalculationResult) error {
	m.published = true
	return m.err
}

// Tests
func TestCalculate_ReturnsCorrectResult(t *testing.T) {
	parser := &mockParser{result: 4}
	publisher := &mockPublisher{}
	service := NewService(parser, publisher)

	result, err := service.Calculate("2+2")

	assert.NoError(t, err)
	assert.Equal(t, 4.0, result.Result)
	assert.Equal(t, "2+2", result.Expression)
}
