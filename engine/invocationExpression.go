package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) InvocationExpression(fhir map[string]interface{}, node grammar.Tree) (interface{}, error) {
	//TODO: call future things
	return []string{"male"}, nil
}
