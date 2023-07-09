package engine

import (
	"errors"
	"fmt"
	"github.com/halprin/fhirpath/internal/grammar"
	"github.com/halprin/rangechain"
	"reflect"
)

func Execute[T any](fhir map[string]interface{}, fhirPathTree grammar.Tree) ([]T, error) {
	fhirOptions := []map[string]interface{}{fhir}
	result, err := (&engine{}).Execute(fhirOptions, fhirPathTree)
	if err != nil {
		return nil, err
	}

	castResult, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("the result of FHIRPath (value=%v, type=%v) cannot be cast into the []interface{} type", result, reflect.TypeOf(result))
	}

	concreteTypeResult, err := filterOutNonRequestedTypes[T](castResult)
	if err != nil {
		return nil, err
	}

	return concreteTypeResult, nil
}

func filterOutNonRequestedTypes[T any](interfaceSlice []interface{}) ([]T, error) {
	filteredInterfaceSlice, err := rangechain.FromSlice(interfaceSlice).Filter(func(currentInterface interface{}) (bool, error) {
		_, ok := currentInterface.(T)
		return ok, nil
	}).Slice()

	if err != nil {
		return nil, err
	}

	filteredRealValues := make([]T, 0, len(filteredInterfaceSlice))
	for _, filteredInterface := range filteredInterfaceSlice {
		realType := filteredInterface.(T)
		filteredRealValues = append(filteredRealValues, realType)
	}

	return filteredRealValues, nil
}

type engine struct {

}

func (receiver *engine) Execute(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {

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

	errorSecondReturn, ok := results[1].Interface().(error)
	if !results[1].IsNil() && !ok {
		return nil, errors.New("engine method " + node.Rule() + " second return value is not an error type")
	}

	return results[0].Interface(), errorSecondReturn
}
