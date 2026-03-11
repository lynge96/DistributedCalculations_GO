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
