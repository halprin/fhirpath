package engine

import (
	"fmt"

	"github.com/halprin/fhirpath/internal/grammar"
)

// IndexerExpression evaluates the left with all of its options, and then returns a specific one depending on the index.
func (receiver *engine) IndexerExpression(fhirOptions []map[string]interface{}, node grammar.Tree) ([]interface{}, error) {
	optionsInterface, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}

	options, ok := optionsInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("IndexerExpression: the left of the index was not a slice (%s)", node.Text())
	}

	indexInterface, err := receiver.Execute(fhirOptions, node.Children()[1])
	if err != nil {
		return nil, err
	}

	index, ok := indexInterface.(int)
	if !ok {
		return nil, fmt.Errorf("IndexerExpression: the index %v is not an integer (%s)", indexInterface, node.Text())
	}

	if len(options)-1 < index {
		return []interface{}{}, nil
	}

	return []interface{}{options[index]}, nil
}
