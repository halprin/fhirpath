package engine

import (
	"errors"
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
			//TODO: except for the `.is(...)` or `.as(...)`.  This conditional is not true, and we need to iterate through the for loop one more time, but we no longer have a FHIR option either.
			break
		}

		accumulator, ok = accumulatorDynamicValue.Value.([]map[string]interface{})
		if !ok {
			//it may still be a slice of FHIR options but hidden behind some stupid Go typing hiding
			interfaceSlice := accumulatorDynamicValue.Value.([]interface{})
			accumulator, ok = convertInterfaceSliceToFhirOptionSlice(interfaceSlice)
			if !ok {
				//we  have more to do but we no longer have a FHIR option, hack it into a FHIR option
				accumulator, err = convertNonFhirOptionToFhirOption(accumulatorDynamicValue)
				if err != nil {
					return nil, errors.New("failed to convert non-FHIR option to FHIR option")
				}
			}
		}
	}

	return accumulatorDynamicValue, nil
}

func convertNonFhirOptionToFhirOption(dynamicValue *DynamicValue) ([]map[string]interface{}, error) {
	sliceSize, err := dynamicValue.SliceSize()
	if err != nil {
		return nil, err
	}

	var fhirOptions []map[string]interface{}

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		currentInterface, err := dynamicValue.SliceValueAtIndex(sliceIndex)
		if err != nil {
			return nil, err
		}

		fhirOptions = append(fhirOptions, map[string]interface{}{
			"value": currentInterface,
		})
	}

	return fhirOptions, nil
}
