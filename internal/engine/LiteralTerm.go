package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// LiteralTerm just evaluates the singular child tree.
func (receiver *engine) LiteralTerm(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return receiver.Execute(fhirOptions, node.Children()[0])
}
