package engine

import (
	"errors"
	"fmt"
	"github.com/halprin/fhirpath/internal/grammar"
)

func (receiver *engine) IndexerExpression(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	optionsInterface, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}
	
	options, ok := optionsInterface.([]interface{})
	if !ok {
		return nil, errors.New("IndexerExpression: the options was not a slice")
	}
	
	indexInterface, err := receiver.Execute(fhirOptions, node.Children()[1])
	if err != nil {
		return nil, err
	}
	
	index, ok := indexInterface.(int)
	if !ok {
		return nil, fmt.Errorf("IndexerExpression: the index %v is not an integer", indexInterface)
	}
	
	return options[index], nil
}
