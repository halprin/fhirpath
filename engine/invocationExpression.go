package engine

import (
	"github.com/halprin/fhirpath/grammar"
	"github.com/halprin/rangechain"
)

func (receiver *engine) InvocationExpression(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return rangechain.FromSlice(node.Children()).ReduceWithInitialValue(func(accumulatorInterface interface{}, currentChildInterface interface{}) (interface{}, error) {
		accumulator, ok := accumulatorInterface.([]map[string]interface{})
		if !ok {
			//it may still be a slice of FHIR options but hidden behind some stupid Go typing hiding
			interfaceSlice := accumulatorInterface.([]interface{})
			accumulator = convertInterfaceSliceToFhirOptionSlice(interfaceSlice)
		}

		currentChild := currentChildInterface.(grammar.Tree)

		return receiver.Execute(accumulator, currentChild)
	}, fhirOptions)
}

//used to convert a generic `[]interface{}` value to a slice of a FHIR option (`map[string]interface{}`)
//this is needed for some of the type casting in the execution engine.  E.g. `InvocationExpression`.
func convertInterfaceSliceToFhirOptionSlice(interfaceSlice []interface{}) []map[string]interface{} {
	fhirOptions := make([]map[string]interface{}, 0, len(interfaceSlice))

	for _, interfaceValue := range interfaceSlice {
		fhirOption := interfaceValue.(map[string]interface{})
		fhirOptions = append(fhirOptions, fhirOption)
	}

	return fhirOptions
}
