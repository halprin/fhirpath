package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// BooleanLiteral converts the string version of the bollean to a boolean.
func (receiver *engine) BooleanLiteral(fhirOptions []map[string]interface{}, node grammar.Tree) (bool, error) {
	literal := node.Text()
	if literal == "true" {
		return true, nil
	}
	return false, nil
}
