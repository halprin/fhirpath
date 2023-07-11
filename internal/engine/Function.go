package engine

import (
	"github.com/halprin/fhirpath/internal/grammar"
)

func (receiver *engine) Function(fhirOptions []map[string]interface{}, node grammar.Tree) ([]interface{}, error) {
	var functionNameAndParams []interface{}
	
	for _, childNode := range node.Children() {
		value, err := receiver.Execute(fhirOptions, childNode)
		if err != nil {
			return nil, err
		}
		
		functionNameAndParams = append(functionNameAndParams, value)
	}
	
	return functionNameAndParams, nil
}
