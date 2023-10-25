package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// StringLiteral strips the start and end single quotes and returns the resulting string.
func (receiver *engine) StringLiteral(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {
	literal := node.Text()
	trimmedLiteral := literal[1 : len(literal)-1] //remove the start and end quotes
	return NewDynamicValue(trimmedLiteral), nil
}
