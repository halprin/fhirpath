package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) Identifier(fhir map[string]interface{}, node grammar.Tree) (interface{}, error) {
	return node.Text(), nil
}
