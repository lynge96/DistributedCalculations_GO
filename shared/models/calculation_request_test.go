package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ReturnsNoError_WhenExpressionIsValid(t *testing.T) {
	request := CalculationRequest{Expression: "2+2"}
	assert.NoError(t, request.Validate())
}

func TestValidate_ReturnsError_WhenExpressionIsEmpty(t *testing.T) {
	request := CalculationRequest{Expression: ""}
	assert.Error(t, request.Validate())
}

func TestValidate_ReturnsError_WhenExpressionExceedsMaxLength(t *testing.T) {
	request := CalculationRequest{Expression: strings.Repeat("a", maxExpressionLength+1)}
	assert.Error(t, request.Validate())
}

func TestValidate_ReturnsNoError_WhenExpressionIsExactlyMaxLength(t *testing.T) {
	request := CalculationRequest{Expression: strings.Repeat("a", maxExpressionLength)}
	assert.NoError(t, request.Validate())
}
