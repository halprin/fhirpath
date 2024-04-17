package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// BooleanLiteral converts the string version of the bollean to a boolean.
func (receiver *engine) BooleanLiteral(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	literal := node.Text()
	if literal == "true" {
		return NewDynamicValue(true), nil
	}
	return NewDynamicValue(false), nil
}
