package fhirpath

import (
	"encoding/json"
	"github.com/halprin/fhirpath/internal/engine"
	"github.com/halprin/fhirpath/internal/grammar"
)

// Evaluate evaluates the FHIR path against the supplied FHIR JSON.
// Returns a slice of values and an optional error.  If the evaluation resulted in nothing, an empty slice is returned.  A slice of size 1 or larger is possible depending on whether the evaluation matched multiple values.
// This function is a generic function, so it takes a type parameter.  Upon evaluation, any results that are not the same as the type parameter are filtered out.  If you want nothing filtered out, use `any` as the type paramter.
func Evaluate[T any](fhirString string, fhirPath string) ([]T, error) {

	fhir, err := unmarshalFhir(fhirString)
	if err != nil {
		return nil, err
	}

	fhir = convertFhirNumbers(fhir)

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

// convertFhirNumbers is a recursive function that converts numbers in the FHIR JSON
// from float64 to int if they are actually integers.
// The function takes a map[string]interface{} as input representing the FHIR JSON.
//
// The function iterates over each key-value pair in the map and calls the helper function
// convertFhirNumbersRecursive to convert any numbers to their corresponding int values.
//
// After iterating through all key-value pairs, the function returns the updated FHIR JSON map.
func convertFhirNumbers(fhir map[string]interface{}) map[string]interface{} {
	for currentKey, currentValue := range fhir {
		fhir[currentKey] = convertFhirNumbersRecursive(currentValue)
	}

	return fhir
}

func convertFhirNumbersRecursive(value interface{}) interface{} {
	switch v := value.(type) {
	case float64:
		// Convert float64 to int if it's really an int
		if float64(int(v)) == v {
			return int(v)
		}
		return v
	case map[string]interface{}:
		// Process map values
		for currentKey, currentValue := range v {
			v[currentKey] = convertFhirNumbersRecursive(currentValue)
		}
	case []interface{}:
		// Process slice values
		for currentIndex, currentValue := range v {
			v[currentIndex] = convertFhirNumbersRecursive(currentValue)
		}
	}
	return value
}
