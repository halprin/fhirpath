package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// ParamList evaluates the children trees in turn and returns the results in turn.
func (receiver *engine) ParamList(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {
	var parameters []interface{}

	for _, childNode := range node.Children() {
		valueDynamicValue, err := receiver.Execute(fhirOptions, childNode)
		if err != nil {
			return nil, err
		}

		parameters = append(parameters, valueDynamicValue.Value)
	}

	return NewDynamicValue(parameters), nil
}
