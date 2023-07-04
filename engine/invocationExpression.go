package engine

import (
	"github.com/halprin/fhirpath/grammar"
	"github.com/halprin/rangechain"
)

func (receiver *engine) InvocationExpression(fhir map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return rangechain.FromSlice(node.Children()).ReduceWithInitialValue(func(accumulatorInterface interface{}, currentChildInterface interface{}) (interface{}, error) {
		accumulator := accumulatorInterface.(map[string]interface{})
		currentChild := currentChildInterface.(grammar.Tree)

		return receiver.Execute(accumulator, currentChild)
	}, fhir)
}
