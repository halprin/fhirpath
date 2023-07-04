package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) TermExpression(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return receiver.Execute(fhirOptions, node.Children()[0])
}
