package engine

import (
	"errors"
	"fmt"

	"github.com/halprin/fhirpath/internal/grammar"
)

// FunctionInvocation evaluates the children and then executes the logic behind the function.
func (receiver *engine) FunctionInvocation(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	functionInterface, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}

	functionConfig, ok := functionInterface.([]interface{})
	if !ok {
		return nil, errors.New("FunctionInvocation: the function configuration was not a slice")
	}

	functionNameInterface := functionConfig[0]
	functionName, ok := functionNameInterface.(string)
	if !ok {
		return nil, errors.New("FunctionInvocation: the function name was not a string")
	}

	functionParameters := []interface{}{}
	if len(functionConfig) > 1 {
		//there are parameters to the function call
		functionParametersInterface := functionConfig[1]
		functionParameters = functionParametersInterface.([]interface{})
		if !ok {
			return nil, errors.New("FunctionInvocation: the function parameters was not a slice")
		}
	}

	//TODO: implement more functions
	switch functionName {
	case "where":
		return where(fhirOptions, functionParameters)
	case "exists":
		return exists(fhirOptions, functionParameters)
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
