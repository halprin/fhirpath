package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

func (receiver *engine) StringLiteral(fhirOptions []map[string]interface{}, node grammar.Tree) (string, error) {
	//remove the start and end quotes
	literal := node.Text()
	trimmedLiteral := literal[1 : len(literal)-1]
	return trimmedLiteral, nil
}
