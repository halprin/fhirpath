package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) TermExpression(fhir map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return receiver.Execute(fhir, node.Children()[0])
}
