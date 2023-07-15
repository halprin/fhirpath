package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// Identifier returns just the grammar.Tree text because it represents a key word.
func (receiver *engine) Identifier(fhirOptions []map[string]interface{}, node grammar.Tree) (string, error) {
	return node.Text(), nil
}
