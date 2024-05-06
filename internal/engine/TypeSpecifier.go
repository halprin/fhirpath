package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// TypeSpecifier specifies a type, like "Integer" or "Quantity".
func (receiver *engine) TypeSpecifier(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	return NewDynamicValue(node.Text()), nil
}
