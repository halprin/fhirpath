package fhirpath

import (
	"encoding/json"
	"github.com/halprin/fhirpath/internal/engine"
	"github.com/halprin/fhirpath/internal/grammar"
)

// Evaluate evaluates the FHIR path against the supplied FHIR JSON.
// Returns a slice of values and an optional error.  If the evaluation resulted in nothing, an empty slice is returned.  A slice of size 1 or larger is possible depending on whether the evaluation matched multiple values.
// This function is a generic function, so it takes a type parameter.  Upon evaluation, any results that are not the same as the type parameter are filtered out.  If you want nothing potentially filtered out, use `any` as the type paramter.
func Evaluate[T any](fhirString string, fhirPath string) ([]T, error) {

	fhir, err := unmarshalFhir(fhirString)
	if err != nil {
		return nil, err
	}

	tree, err := grammar.CreateTree(fhirPath)
	if err != nil {
		return nil, err
	}

	result, err := engine.Execute[T](fhir, tree)
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
