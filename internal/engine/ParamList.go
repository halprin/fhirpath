package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

// ParamList evaluates the children trees in turn and returns the results in turn.
func (receiver *engine) ParamList(fhirOptions []map[string]interface{}, node grammar.Tree) ([]interface{}, error) {
	var parameters []interface{}

	for _, childNode := range node.Children() {
		value, err := receiver.Execute(fhirOptions, childNode)
		if err != nil {
			return nil, err
		}

		parameters = append(parameters, value)
	}

	return parameters, nil
}
