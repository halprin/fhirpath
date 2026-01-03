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
		filteredValues, err := asOperation(leftOperands, rightOperands)
		if err != nil {
			return nil, err
		}

		return NewDynamicValue(filteredValues), nil
	}

	return nil, fmt.Errorf("TypeExpression is not 'is' or 'as', instead it is %s", operation)
}

func asOperation(dynamicValue *DynamicValue, dynamicTypeIdentifier *DynamicValue) ([]map[string]interface{}, error) {
	// Check which values match the type
	isResults, err := isOperation(dynamicValue, dynamicTypeIdentifier)
	if err != nil {
		return nil, err
	}

	sliceSize, err := dynamicValue.SliceSize()
	if err != nil {
		return nil, err
	}

	// Return only values that match the type
	var filteredValues []map[string]interface{}
	for i := 0; i < sliceSize; i++ {
		if i < len(isResults) && isResults[i] {
			value, err := CastSliceValueAtIndexOfDynamicValue[map[string]interface{}](dynamicValue, i)
			//value, err := dynamicValue.SliceValueAtIndex(i)
			if err != nil {
				return nil, err
			}
			filteredValues = append(filteredValues, value)
		}
	}

	return filteredValues, nil
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
		isTypeSlice, err = isDynamicValueSliceIsQuantity(dynamicValue)
	}

	//TODO: implement the FHIR types

	return isTypeSlice, err
}

func isDynamicValueSliceIsQuantity(dynamicValue *DynamicValue) ([]bool, error) {
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

		// Quantity is a map with a "value" key containing a number
		mapValue, ok := currentInterface.(map[string]interface{})
		if !ok {
			isTypeSlice = append(isTypeSlice, false)
			continue
		}

		// Check if it has a "value" key with a numeric value (required for Quantity)
		quantityValue, hasValue := mapValue["value"]
		if !hasValue {
			isTypeSlice = append(isTypeSlice, false)
			continue
		}

		// Value should be numeric (float64 from JSON)
		_, isFloat := quantityValue.(float64)
		_, isInt := quantityValue.(int)
		isTypeSlice = append(isTypeSlice, isFloat || isInt)
	}

	return isTypeSlice, nil
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
