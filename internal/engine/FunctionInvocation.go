package engine

import (
	"errors"
	"fmt"

	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
)

// FunctionInvocation evaluates the children and then executes the logic behind the function.
func (receiver *engine) FunctionInvocation(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	functionDynamicValue, err := receiver.Execute(fhirOptions, node.Children()[0], context)
	if err != nil {
		return nil, err
	}

	functionName, err := CastSliceValueAtIndexOfDynamicValue[string](functionDynamicValue, 0)
	if err != nil {
		return nil, errors.Join(errors.New("FunctionInvocation: function name failure"), err)
	}

	functionParameters := []interface{}{}

	functionConfigSize, err := functionDynamicValue.SliceSize()
	if err != nil {
		return nil, errors.Join(errors.New("FunctionInvocation: function parameters failure"), err)
	}

	if functionConfigSize > 1 {
		//there are parameters to the function call
		functionParameters, err = CastSliceValueAtIndexOfDynamicValue[[]interface{}](functionDynamicValue, 1)
		if err != nil {
			return nil, errors.Join(errors.New("FunctionInvocation: function parameters failure"), err)
		}
	}

	//TODO: implement more functions
	switch functionName {
	case "where":
		newFhirOptions, err := where(fhirOptions, functionParameters)
		return NewDynamicValue(newFhirOptions), err
	case "exists":
		bools, err := exists(fhirOptions, functionParameters)
		return NewDynamicValue(bools), err
	case "is":
		bools, err := is(fhirOptions, functionParameters)
		return NewDynamicValue(bools), err
	case "as":
		filteredOptions, err := as(fhirOptions, functionParameters)
		return NewDynamicValue(filteredOptions), err
	case "empty":
		return NewDynamicValue([]bool{len(fhirOptions) == 0}), nil
	case "not":
		bools, err := not(fhirOptions)
		return NewDynamicValue(bools), err
	case "first":
		return NewDynamicValue(first(fhirOptions)), nil
	case "last":
		return NewDynamicValue(last(fhirOptions)), nil
	case "count":
		return NewDynamicValue([]int{len(fhirOptions)}), nil
	case "skip":
		skipped, err := skip(fhirOptions, functionParameters)
		return NewDynamicValue(skipped), err
	case "take":
		taken, err := take(fhirOptions, functionParameters)
		return NewDynamicValue(taken), err
	default:
		return nil, fmt.Errorf("FunctionInvocation: function name %s is unknown", functionName)
	}
}

func where(fhirOptions []map[string]interface{}, parameters []interface{}) ([]map[string]interface{}, error) {
	//where has only 1 parameter: the evaluation of the expression inside

	booleanEvaluation, ok := parameters[0].([]bool)
	if !ok {
		return nil, errors.New("FunctionInvocation: where: the first parameter was not a boolean slice")
	}

	var filteredFhirOptions []map[string]interface{}

	for index, currentFhirOption := range fhirOptions {
		if !booleanEvaluation[index] {
			//the evaluation was found not matching so we filter OUT this FHIR option
			continue
		}

		filteredFhirOptions = append(filteredFhirOptions, currentFhirOption)
	}

	return filteredFhirOptions, nil
}

func exists(fhirOptions []map[string]interface{}, parameters []interface{}) ([]bool, error) {
	if len(parameters) > 0 {
		//there were parameters which is the equivalent of running where first
		var err error
		fhirOptions, err = where(fhirOptions, parameters)
		if err != nil {
			return nil, err
		}
	}

	if len(fhirOptions) == 0 {
		return []bool{false}, nil
	}

	return []bool{true}, nil
}

func is(fhirOptions []map[string]interface{}, parameters []interface{}) ([]bool, error) {
	if len(parameters) == 0 {
		return nil, errors.New("FunctionInvocation: is: requires a type parameter")
	}

	typeIdentifier, ok := parameters[0].(string)
	if !ok {
		return nil, errors.New("FunctionInvocation: is: the type parameter must be a string")
	}

	if len(fhirOptions) == 0 {
		return []bool{}, nil
	}

	// Extract actual values from fhirOptions
	// Primitive values are wrapped in {"value": primitive} by convertNonFhirOptionToFhirOption (single key)
	// Complex FHIR types like Quantity have a "value" field plus other fields (multiple keys)
	var values []interface{}
	for _, fhirOption := range fhirOptions {
		if len(fhirOption) == 1 {
			if value, hasValue := fhirOption["value"]; hasValue {
				// Single "value" key - this is a wrapped primitive
				values = append(values, value)
				continue
			}
		}
		// It's a full FHIR complex type, add it as-is
		values = append(values, fhirOption)
	}

	dynamicValue := NewDynamicValue(values)
	dynamicTypeIdentifier := NewDynamicValue(typeIdentifier)

	return isOperation(dynamicValue, dynamicTypeIdentifier)
}

func as(fhirOptions []map[string]interface{}, parameters []interface{}) ([]map[string]interface{}, error) {
	if len(parameters) == 0 {
		return nil, errors.New("FunctionInvocation: as: requires a type parameter")
	}

	typeIdentifier, ok := parameters[0].(string)
	if !ok {
		return nil, errors.New("FunctionInvocation: as: the type parameter must be a string")
	}

	if len(fhirOptions) == 0 {
		return []map[string]interface{}{}, nil
	}

	dynamicValue := NewDynamicValue(fhirOptions)
	dynamicTypeIdentifier := NewDynamicValue(typeIdentifier)

	return asOperation(dynamicValue, dynamicTypeIdentifier)
}

func not(fhirOptions []map[string]interface{}) ([]bool, error) {
	if len(fhirOptions) == 0 {
		return []bool{}, nil
	}

	// Extract the boolean value from fhirOptions
	// The value should be wrapped in {"value": bool}
	if len(fhirOptions) != 1 {
		return nil, errors.New("FunctionInvocation: not: expected a single boolean value")
	}

	value, ok := fhirOptions[0]["value"]
	if !ok {
		return nil, errors.New("FunctionInvocation: not: expected a wrapped boolean value")
	}

	boolValue, ok := value.(bool)
	if !ok {
		return nil, errors.New("FunctionInvocation: not: value is not a boolean")
	}

	return []bool{!boolValue}, nil
}

func first(fhirOptions []map[string]interface{}) []map[string]interface{} {
	if len(fhirOptions) == 0 {
		return []map[string]interface{}{}
	}
	return []map[string]interface{}{fhirOptions[0]}
}

func last(fhirOptions []map[string]interface{}) []map[string]interface{} {
	if len(fhirOptions) == 0 {
		return []map[string]interface{}{}
	}
	return []map[string]interface{}{fhirOptions[len(fhirOptions)-1]}
}

func skip(fhirOptions []map[string]interface{}, parameters []interface{}) ([]map[string]interface{}, error) {
	if len(parameters) == 0 {
		return nil, errors.New("FunctionInvocation: skip: requires a count parameter")
	}

	// The parameter could be an int or wrapped in a slice
	var count int
	switch v := parameters[0].(type) {
	case int:
		count = v
	case float64:
		count = int(v)
	case []int:
		if len(v) > 0 {
			count = v[0]
		}
	default:
		return nil, errors.New("FunctionInvocation: skip: count parameter must be an integer")
	}

	if count < 0 {
		count = 0
	}
	if count >= len(fhirOptions) {
		return []map[string]interface{}{}, nil
	}

	return fhirOptions[count:], nil
}

func take(fhirOptions []map[string]interface{}, parameters []interface{}) ([]map[string]interface{}, error) {
	if len(parameters) == 0 {
		return nil, errors.New("FunctionInvocation: take: requires a count parameter")
	}

	// The parameter could be an int or wrapped in a slice
	var count int
	switch v := parameters[0].(type) {
	case int:
		count = v
	case float64:
		count = int(v)
	case []int:
		if len(v) > 0 {
			count = v[0]
		}
	default:
		return nil, errors.New("FunctionInvocation: take: count parameter must be an integer")
	}

	if count < 0 {
		count = 0
	}
	if count >= len(fhirOptions) {
		return fhirOptions, nil
	}

	return fhirOptions[:count], nil
}
