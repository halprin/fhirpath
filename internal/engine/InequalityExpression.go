package engine

import (
	"fmt"

	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// InequalityExpression evaluates comparison operations: <, >, <=, >=
func (receiver *engine) InequalityExpression(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	leftOperands := receiver.populateOperands(fhirOptions, node.Children()[0], context)
	rightOperands := receiver.populateOperands(fhirOptions, node.Children()[1], context)

	operation := node.TerminalTexts()[0]

	switch operation {
	case "<":
		return compareSlices(fhirOptions, leftOperands, rightOperands, lessThan)
	case ">":
		return compareSlices(fhirOptions, leftOperands, rightOperands, greaterThan)
	case "<=":
		return compareSlices(fhirOptions, leftOperands, rightOperands, lessThanOrEqual)
	case ">=":
		return compareSlices(fhirOptions, leftOperands, rightOperands, greaterThanOrEqual)
	default:
		return nil, fmt.Errorf("InequalityExpression: operation %s is unknown", operation)
	}
}

func lessThan(leftOperand interface{}, rightOperand interface{}) bool {
	leftFloat, rightFloat, ok := toComparableFloats(leftOperand, rightOperand)
	if !ok {
		return false
	}
	return leftFloat < rightFloat
}

func greaterThan(leftOperand interface{}, rightOperand interface{}) bool {
	leftFloat, rightFloat, ok := toComparableFloats(leftOperand, rightOperand)
	if !ok {
		return false
	}
	return leftFloat > rightFloat
}

func lessThanOrEqual(leftOperand interface{}, rightOperand interface{}) bool {
	leftFloat, rightFloat, ok := toComparableFloats(leftOperand, rightOperand)
	if !ok {
		return false
	}
	return leftFloat <= rightFloat
}

func greaterThanOrEqual(leftOperand interface{}, rightOperand interface{}) bool {
	leftFloat, rightFloat, ok := toComparableFloats(leftOperand, rightOperand)
	if !ok {
		return false
	}
	return leftFloat >= rightFloat
}

func toComparableFloats(left interface{}, right interface{}) (float64, float64, bool) {
	leftFloat, ok := toFloat64(left)
	if !ok {
		return 0, 0, false
	}

	rightFloat, ok := toFloat64(right)
	if !ok {
		return 0, 0, false
	}

	return leftFloat, rightFloat, true
}

func toFloat64(val interface{}) (float64, bool) {
	switch v := val.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	default:
		return 0, false
	}
}
