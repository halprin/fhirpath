// Package engine evaluates the grammar.Tree and returns the results.
package engine

import (
	"errors"
	"reflect"

	"github.com/halprin/fhirpath/internal/grammar"
)

// Execute starts the evaluation of the grammar.Tree given the FHIR object.
// It also filters out types that don't match the type parameter.
func Execute[T any](fhir map[string]interface{}, fhirPathTree grammar.Tree) ([]T, error) {
	fhirOptions := []map[string]interface{}{fhir}
	result, err := (&engine{}).Execute(fhirOptions, fhirPathTree)
	if err != nil {
		return nil, err
	}

	concreteTypeResult, err := CastAndFilterSliceOfDynamicValue[T](result)
	if err != nil {
		return nil, err
	}

	return concreteTypeResult, nil
}

type engine struct {
}

// Execute dynamically calls the engine's method that matches the rule of the current grammar.Tree.
func (receiver *engine) Execute(fhirOptions []map[string]interface{}, node grammar.Tree) (*DynamicValue, error) {

	engineReflect := reflect.ValueOf(receiver)

	engineMethod := engineReflect.MethodByName(node.Rule())
	if !engineMethod.IsValid() {
		return nil, errors.New("engine method " + node.Rule() + " doesn't exist")
	}

	methodArguments := []reflect.Value{reflect.ValueOf(fhirOptions), reflect.ValueOf(node)}
	results := engineMethod.Call(methodArguments)
	if len(results) != 2 {
		return nil, errors.New("engine method " + node.Rule() + " doesn't return two values")
	}

	valueFirstReturn, ok := results[0].Interface().(*DynamicValue)
	if !results[0].IsNil() && !ok {
		return nil, errors.New("engine method " + node.Rule() + " second return value is not an DynamicValue pointer type")
	}

	errorSecondReturn, ok := results[1].Interface().(error)
	if !results[1].IsNil() && !ok {
		return nil, errors.New("engine method " + node.Rule() + " second return value is not an error type")
	}

	return valueFirstReturn, errorSecondReturn
}
