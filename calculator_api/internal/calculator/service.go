package calculator

import (
	"calculator_api/internal/models"
	"log"
	"time"

	"github.com/google/uuid"
)

type MathParser interface {
	Evaluate(string) (float64, error)
}

type Service struct {
	parser MathParser
}

func NewService(parser MathParser) *Service {
	return &Service{parser: parser}
}

func (s *Service) Calculate(expression string) (models.CalculationResult, error) {

	result, err := s.parser.Evaluate(expression)
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
