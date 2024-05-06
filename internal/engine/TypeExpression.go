package engine

import (
	"fmt"
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// TypeExpression executes operations like "is" or "as".
func (receiver *engine) TypeExpression(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {

	leftOperands, err := receiver.Execute(fhirOptions, node.Children()[0], context)
	if err != nil {
		return nil, err
	}

	rightOperands, err := receiver.Execute(fhirOptions, node.Children()[1], context)
	if err != nil {
		return nil, err
	}

	operation := node.TerminalTexts()[0]
	fmt.Println(operation, leftOperands, rightOperands)
	return receiver.Execute(fhirOptions, node.Children()[0], context)
}
