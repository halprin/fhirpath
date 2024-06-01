package engine

import (
	"errors"
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

	if operation == "is" {
		isTypeSlice, err := isOperation(leftOperands, rightOperands)
		if err != nil {
			return nil, err
		}

		return NewDynamicValue(isTypeSlice), nil
	} else if operation == "as" {
		return nil, errors.New("TypeExpression 'as' needs to be implemented")
	} else {
		return nil, fmt.Errorf("TypeExpression is not 'is' or 'as', instead it is %s", operation)
	}
}

func isOperation(dynamicValue *DynamicValue, dynamicTypeIdentifier *DynamicValue) ([]bool, error) {
	typeIdentifier, ok := dynamicTypeIdentifier.Value.(string)
	if !ok {
		return nil, errors.New("the type identifier in a TypeExpression is not a string")
	}

	var isTypeSlice []bool
	var err error

	//TODO: implement more literal types
	switch typeIdentifier {
	case "Boolean":
		isTypeSlice, err = isDynamicValueSliceIsType[bool](dynamicValue)
	case "String":
		isTypeSlice, err = isDynamicValueSliceIsType[string](dynamicValue)
	case "Integer":
		isTypeSlice, err = isDynamicValueSliceIsType[int](dynamicValue)
	case "Decimal":
		isTypeSlice, err = isDynamicValueSliceIsType[float64](dynamicValue)
	case "Date":
		return nil, errors.New("TypeExpression doesn't support the is operation with the Date type yet, it needs to be implemented")
	case "DateTime":
		return nil, errors.New("TypeExpression doesn't support the is operation with the DateTime type yet, it needs to be implemented")
	case "Time":
		return nil, errors.New("TypeExpression doesn't support the is operation with the Time type yet, it needs to be implemented")
	case "Quantity":
		return nil, errors.New("TypeExpression doesn't support the is operation with the Quantity type yet, it needs to be implemented")
	}

	//TODO: implement the FHIR types

	return isTypeSlice, err
}

func isDynamicValueSliceIsType[T any](dynamicValue *DynamicValue) ([]bool, error) {
	sliceSize, err := dynamicValue.SliceSize()
	if err != nil {
		return nil, err
	}

	var isTypeSlice []bool

	for sliceIndex := 0; sliceIndex < sliceSize; sliceIndex++ {
		currentInterface, err := dynamicValue.SliceValueAtIndex(sliceIndex)
		if err != nil {
			return nil, err
		}

		_, ok := currentInterface.(T)
		isTypeSlice = append(isTypeSlice, ok)
	}

	return isTypeSlice, nil
}
