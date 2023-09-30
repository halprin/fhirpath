package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// TermExpression just evaluates the singular child tree.
func (receiver *engine) TermExpression(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {
	return receiver.Execute(fhirOptions, node.Children()[0])
}
