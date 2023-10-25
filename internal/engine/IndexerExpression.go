package engine

import (
	"errors"
	"fmt"

	"github.com/halprin/fhirpath/internal/grammar"
)

// IndexerExpression evaluates the left with all of its options, and then returns a specific one depending on the index.
func (receiver *engine) IndexerExpression(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {
	optionsDynamicValue, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}

	indexDynamicValue, err := receiver.Execute(fhirOptions, node.Children()[1])
	if err != nil {
		return nil, err
	}

	index, ok := indexDynamicValue.Value.(int)
	if !ok {
		return nil, fmt.Errorf("IndexerExpression: the index %v is not an integer (%s)", indexDynamicValue, node.Text())
	}

	optionsSize, err := optionsDynamicValue.SliceSize()
	if err != nil {
		return nil, errors.Join(errors.New("IndexerExpression: failure determining the options size"), err)
	}

	if optionsSize-1 < index {
		return NewDynamicValue([]interface{}{}), nil
	}

	valueAtIndex, err := optionsDynamicValue.SliceValueAtIndex(index)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("IndexerExpression: failure getting the value at index %d", index), err)
	}

	return NewDynamicValue([]interface{}{valueAtIndex}), nil
}
