package models

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCalculationResult_SetsExpressionAndResult(t *testing.T) {
	result := NewCalculationResult("2+2", 4)
	assert.Equal(t, "2+2", result.Expression)
	assert.Equal(t, 4.0, result.Result)
}

func TestNewCalculationResult_SetsIDAndTimeStamp(t *testing.T) {
	result := NewCalculationResult("2+2", 4)
	assert.NotEqual(t, uuid.Nil, result.ID)
	assert.NotEqual(t, 0, result.Timestamp.Unix())
}

func TestNewCalculationError_SetsError(t *testing.T) {
	result := NewCalculationError("invalid", fmt.Errorf("unknown operation"))
	assert.Equal(t, "unknown operation", result.Error)
}
