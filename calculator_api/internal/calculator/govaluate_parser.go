package calculator

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/Knetic/govaluate"
)

type GovaluateParser struct{}

var functions = map[string]govaluate.ExpressionFunction{
	"SIN": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("SIN expects 1 argument")
		}
		return math.Sin(args[0].(float64)), nil
	},
	"COS": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("COS expects 1 argument")
		}
		return math.Cos(args[0].(float64)), nil
	},
	"TAN": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("TAN expects 1 argument")
		}
		return math.Tan(args[0].(float64)), nil
	},
	"SQRT": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("SQRT expects 1 argument")
		}
		return math.Sqrt(args[0].(float64)), nil
	},
	"ABS": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("ABS expects 1 argument")
		}
		return math.Abs(args[0].(float64)), nil
	},
	"LOG": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("LOG expects 1 argument")
		}
		return math.Log10(args[0].(float64)), nil
	},
}

func (p *GovaluateParser) Evaluate(expression string) (float64, error) {

	expression = strings.ReplaceAll(expression, "^", "**")

	exp, err := govaluate.NewEvaluableExpressionWithFunctions(expression, functions)
	if err != nil {
		return 0, err
	}
	log.Printf("Evaluating expression: %v", exp)

	parameters := map[string]interface{}{
		"PI": math.Pi,
	}

	result, err := exp.Evaluate(parameters)
	if err != nil {
		return 0, err
	}

	log.Printf("Result: %+v", result)
	return result.(float64), nil
}
