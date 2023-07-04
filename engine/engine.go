package engine

import (
	"errors"
	"github.com/halprin/fhirpath/grammar"
	"reflect"
)

func Execute[T any](fhir map[string]interface{}, fhirPathTree grammar.Tree) ([]T, error) {
	fhirOptions := []map[string]interface{}{fhir}
	result, err := (&engine{}).Execute(fhirOptions, fhirPathTree)
	if err != nil {
		return nil, err
	}

	castResult, ok := result.([]T)
	if !ok {
		return nil, errors.New("the result of FHIRPath cannot be cast into the requested type")
	}

	return castResult, nil
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
