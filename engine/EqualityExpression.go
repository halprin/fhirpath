package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) EqualityExpression(fhirOptions []map[string]interface{}, node grammar.Tree) ([]bool, error) {
	//evaluate each operand against each FHIR option
	
	leftOperands := receiver.populateOperands(fhirOptions, node.Children()[0])
	rightOperands := receiver.populateOperands(fhirOptions, node.Children()[1])
	
	//TODO: there are multiple terminal text that we need to parse out.  E.g. `=`, `!=`, etc.
	
	equalitySlice := make([]bool, len(fhirOptions))
	
	for index, _ := range fhirOptions {
		equalitySlice[index] = leftOperands[index] == rightOperands[index]
	}
	
	return equalitySlice, nil
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
