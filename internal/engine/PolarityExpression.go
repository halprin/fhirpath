package engine

import (
	"errors"

	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// PolarityExpression handles unary + and - operations on numeric values
func (receiver *engine) PolarityExpression(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	operand, err := receiver.Execute(fhirOptions, node.Children()[0], context)
	if err != nil {
		return nil, err
	}

	operation := node.TerminalTexts()[0]

	// Handle the case where operand is a slice
	if operand.IsSlice() {
		sliceSize, err := operand.SliceSize()
		if err != nil {
			return nil, err
		}

		var results []interface{}
		for i := 0; i < sliceSize; i++ {
			val, err := operand.SliceValueAtIndex(i)
			if err != nil {
				return nil, err
			}

			result, err := applyPolarity(val, operation)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
		return NewDynamicValue(results), nil
	}

	// Handle single value
	result, err := applyPolarity(operand.Value, operation)
	if err != nil {
		return nil, err
	}
	return NewDynamicValue(result), nil
}

func applyPolarity(val interface{}, operation string) (interface{}, error) {
	switch v := val.(type) {
	case int:
		if operation == "-" {
			return -v, nil
		}
		return v, nil
	case float64:
		if operation == "-" {
			return -v, nil
		}
		return v, nil
	case int64:
		if operation == "-" {
			return -v, nil
		}
		return v, nil
	default:
		return nil, errors.New("PolarityExpression: operand is not a numeric type")
	}
}
