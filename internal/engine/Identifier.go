package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// Identifier returns the grammar.Tree text because it represents a key word.
// If the identifier is enclosed in backticks, the backticks are removed.
func (receiver *engine) Identifier(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	text := node.Text()

	if len(text) >= 2 && text[0] == '`' && text[len(text)-1] == '`' {
		// Remove the backticks
		text = text[1 : len(text)-1]
	}

	return NewDynamicValue(text), nil
}
