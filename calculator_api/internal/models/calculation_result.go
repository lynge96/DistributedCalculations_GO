package models

import (
	"time"

	"github.com/google/uuid"
)

type CalculationResult struct {
	ID         uuid.UUID `json:"id"`
	Expression string    `json:"expression"`
	Result     float64   `json:"result"`
	Timestamp  time.Time `json:"timestamp"`
}
