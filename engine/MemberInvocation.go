package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) MemberInvocation(fhir map[string]interface{}, node grammar.Tree) (interface{}, error) {
	key, err := receiver.Execute(fhir, node.Children()[0])
	if err != nil {
		return nil, err
	}
	
	return key, nil
}
