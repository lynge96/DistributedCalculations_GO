package models

import (
	"fmt"
	"strings"
)

const maxExpressionLength = 256

type CalculationRequest struct {
	Expression string `json:"expression"`
}

func (r CalculationRequest) Validate() error {
	if strings.TrimSpace(r.Expression) == "" {
		return fmt.Errorf("expression cannot be empty")
	}
	if len(r.Expression) > maxExpressionLength {
		return fmt.Errorf("expression cannot be longer than %d characters", maxExpressionLength)
	}
	return nil
}
