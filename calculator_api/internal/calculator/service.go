package calculator

import (
	"log"
	"shared/models"
	"time"

	"github.com/google/uuid"
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
		return models.CalculationResult{}, err
	}

	calculationResult := models.CalculationResult{
		ID:         uuid.New(),
		Expression: expression,
		Result:     result,
		Timestamp:  time.Now(),
	}

	if err := s.publisher.Publish(calculationResult); err != nil {
		log.Printf("failed to publish result: %v", err)
	}

	log.Printf("Calculation result: %v", calculationResult)
	return calculationResult, nil
}
