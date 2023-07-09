package engine

import (
	"github.com/halprin/fhirpath/grammar"
)

func (receiver *engine) FunctionInvocation(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	args, err := receiver.Execute(fhirOptions, node.Children()[0])
	//TODO: apply the stuff learned by your children to do the function.  E.g. `where`.  Or could this be done in `Function`?
	return args, err
}
