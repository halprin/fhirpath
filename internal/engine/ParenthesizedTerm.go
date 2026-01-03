package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// ParenthesizedTerm evaluates the expression inside the parentheses.
func (receiver *engine) ParenthesizedTerm(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	return receiver.Execute(fhirOptions, node.Children()[0], context)
}
