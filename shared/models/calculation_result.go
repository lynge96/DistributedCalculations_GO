package models

import (
	"time"

	"github.com/google/uuid"
)

type CalculationResult struct {
	ID         uuid.UUID `json:"id"`
	Expression string    `json:"expression,omitempty"`
	Result     float64   `json:"result,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	Error      string    `json:"error,omitempty"`
}

func NewCalculationResult(expression string, result float64) CalculationResult {
	return CalculationResult{
		ID:         uuid.New(),
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now(),
	}
}

func NewCalculationError(expression string, err error) CalculationResult {
	return CalculationResult{
		ID:         uuid.New(),
		Expression: expression,
		Error:      err.Error(),
		Timestamp:  time.Now(),
	}
}
