package calculator

import (
	"github.com/Knetic/govaluate"
)

func Evaluate(expression string) (float64, error) {
	
	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	result, err := exp.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}
