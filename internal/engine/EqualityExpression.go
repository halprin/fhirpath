package engine

import (
	"fmt"

	"github.com/halprin/fhirpath/internal/grammar"
)

// EqualityExpression evaluates each child node against each FHIR option.  It then does the operation (equaliity, equivalent, not) between the right and the left, for each FHIR option.  This slice of boolean values is returned.
func (receiver *engine) EqualityExpression(fhirOptions []map[string]interface{}, node grammar.Tree) ([]bool, error) {
	//evaluate each operand against each FHIR option

	leftOperands := receiver.populateOperands(fhirOptions, node.Children()[0])
	rightOperands := receiver.populateOperands(fhirOptions, node.Children()[1])

	operation := node.TerminalTexts()[0]

	//TODO: implement equivalent and not equivalent.  https://hl7.org/fhirpath/#equivalent
	switch operation {
	case "=":
		return compareSlices(fhirOptions, leftOperands, rightOperands, equals)
	case "!=":
		return compareSlices(fhirOptions, leftOperands, rightOperands, notEquals)
	default:
		return nil, fmt.Errorf("EqualityExpression: operation %s is unknown", operation)
	}
}

func (receiver *engine) populateOperands(fhirOptions []map[string]interface{}, node grammar.Tree) []interface{} {
	var operandValues []interface{}

	for _, fhirOption := range fhirOptions {
		wrapInSlice := []map[string]interface{}{fhirOption}
		leftOperand, err := receiver.Execute(wrapInSlice, node)
		if err != nil {
			return nil
		}

		operandValues = append(operandValues, leftOperand)
	}

	operandValues = flatten(operandValues)

	return operandValues
}

func compareSlices(fhirOptions []map[string]interface{}, leftOperands []interface{}, rightOperands []interface{}, comparisonFunction func(interface{}, interface{}) bool) ([]bool, error) {
	comparisonSlice := make([]bool, len(fhirOptions))

	for index, _ := range fhirOptions {
		comparisonSlice[index] = comparisonFunction(leftOperands[index], rightOperands[index])
	}

	return comparisonSlice, nil
}

func equals(leftOperand interface{}, rightOperand interface{}) bool {
	return leftOperand == rightOperand
}

func notEquals(leftOperand interface{}, rightOperand interface{}) bool {
	return leftOperand != rightOperand
}
