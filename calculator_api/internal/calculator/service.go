package calculator

import (
	"log"
	"shared/models"
)

type MathParser interface {
	Evaluate(string) (float64, error)
}

type Publisher interface {
	Publish(message models.CalculationResult) error
}

type Service struct {
	parser    MathParser
	publisher Publisher
}

func NewService(parser MathParser, publisher Publisher) *Service {
	return &Service{
		parser:    parser,
		publisher: publisher,
	}
}

func (s *Service) Calculate(expression string) (models.CalculationResult, error) {

	result, err := s.parser.Evaluate(expression)
	if err != nil {
		return models.NewCalculationError(expression, err), nil
	}

	calculationResult := models.NewCalculationResult(expression, result)

	if err := s.publisher.Publish(calculationResult); err != nil {
		log.Printf("failed to publish result: %v", err)
	}

	log.Printf("Calculation result: %v", calculationResult)
	return calculationResult, nil
}
