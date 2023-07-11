package engine

import (
	"fmt"
	"github.com/halprin/fhirpath/internal/grammar"
)

func (receiver *engine) EqualityExpression(fhirOptions []map[string]interface{}, node grammar.Tree) ([]bool, error) {
	//evaluate each operand against each FHIR option
	
	leftOperands := receiver.populateOperands(fhirOptions, node.Children()[0])
	rightOperands := receiver.populateOperands(fhirOptions, node.Children()[1])
	

	operation := node.TerminalTexts()[0]

	//TODO: implement more operations
	switch operation {
	case "=":
		return equals(fhirOptions, leftOperands, rightOperands)
	case "!=":
		return notEquals(fhirOptions, leftOperands, rightOperands)
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

func equals(fhirOptions []map[string]interface{}, leftOperands []interface{}, rightOperands []interface{}) ([]bool, error) {
	equalityFunction := func(leftOperand interface{}, rightOperand interface{}) bool {
		return leftOperand == rightOperand
	}

	return compareSlices(fhirOptions, leftOperands, rightOperands, equalityFunction)
}

func notEquals(fhirOptions []map[string]interface{}, leftOperands []interface{}, rightOperands []interface{}) ([]bool, error) {
	notEqualityFunction := func(leftOperand interface{}, rightOperand interface{}) bool {
		return leftOperand != rightOperand
	}

	return compareSlices(fhirOptions, leftOperands, rightOperands, notEqualityFunction)
}

func compareSlices(fhirOptions []map[string]interface{}, leftOperands []interface{}, rightOperands []interface{}, comparisonFunction func(interface{}, interface{}) bool) ([]bool, error) {
	comparisonSlice := make([]bool, len(fhirOptions))

	for index, _ := range fhirOptions {
		comparisonSlice[index] = comparisonFunction(leftOperands[index], rightOperands[index])
	}

	return comparisonSlice, nil
}
