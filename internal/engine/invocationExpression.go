package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
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
