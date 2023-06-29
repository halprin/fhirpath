package fhirpath

import (
	"encoding/json"
	"github.com/halprin/fhirpath/grammar"
)

func Evaluate[T any](fhirString string, fhirPath string) ([]T, error) {
	
	fhir, err := unmarshalFhir(fhirString)
	if err != nil {
		return nil, err
	}

	result, err := grammar.AntlrExecute[T](fhir, fhirPath)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func unmarshalFhir(fhir string) (map[string]interface{}, error) {
	var fhirObject map[string]interface{}

	err := json.Unmarshal([]byte(fhir), &fhirObject)
	if err != nil {
		return nil, err
	}

	return fhirObject, nil
}
