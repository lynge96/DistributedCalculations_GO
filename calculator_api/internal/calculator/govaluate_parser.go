package calculator

import (
	"log"

	"github.com/Knetic/govaluate"
)

type GovaluateParser struct{}

func (p *GovaluateParser) Evaluate(expression string) (float64, error) {

	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}
	log.Printf("Evaluating expression: %v", exp)

	result, err := exp.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	log.Printf("Result: %v", result)
	return result.(float64), nil
}
