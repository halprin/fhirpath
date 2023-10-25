package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// InvocationTerm just evaluates the singular child tree.
func (receiver *engine) InvocationTerm(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {
	return receiver.Execute(fhirOptions, node.Children()[0])
}
