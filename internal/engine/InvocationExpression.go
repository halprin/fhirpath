package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// InvocationExpression is all about evaluating a child tree, and then evaluate the next child tree given any FHIR option results that came from the previous tree evaluation.
func (receiver *engine) InvocationExpression(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {

	accumulator := fhirOptions
	var accumulatorDynamicValue *DynamicValue
	var err error
	ok := false

	for index, currentChild := range node.Children() {
		accumulatorDynamicValue, err = receiver.Execute(accumulator, currentChild, context)
		if err != nil {
			return nil, err
		}

		if index == len(node.Children())-1 {
			//this is the last iteration
			//break early because the last iteration probably returned something different than a FHIR option, so don't try to convert to a FHIR option
			break
		}

		accumulator, ok = accumulatorDynamicValue.Value.([]map[string]interface{})
		if !ok {
			//it may still be a slice of FHIR options but hidden behind some stupid Go typing hiding
			interfaceSlice := accumulatorDynamicValue.Value.([]interface{})
			accumulator = convertInterfaceSliceToFhirOptionSlice(interfaceSlice)
		}
	}

	return accumulatorDynamicValue, nil
}
