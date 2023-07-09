package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) FunctionInvocation(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	args, err := receiver.Execute(fhirOptions, node.Children()[0])
	return args, err
}
