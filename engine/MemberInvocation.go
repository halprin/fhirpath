package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) MemberInvocation(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	key, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}
	
	return key, nil
}
