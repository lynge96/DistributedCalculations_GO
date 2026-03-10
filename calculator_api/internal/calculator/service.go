package calculator

import (
	"calculator_api/pkg/models"
	"log"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	logger log.Logger
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(expression string) (models.CalculationResult, error) {

	result, err := Evaluate(expression)
	if err != nil {
		return models.CalculationResult{}, err
	}

	calculationResult := models.CalculationResult{
		ID:         uuid.New(),
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now(),
	}

	log.Printf("Calculation result: %v", calculationResult)
	return calculationResult, nil
}
