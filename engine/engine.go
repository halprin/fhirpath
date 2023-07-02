package engine

import "github.com/halprin/fhirpath/grammar"

func Execute[T any](fhir map[string]interface{}, fhirPathTree grammar.Tree) ([]T, error) {
	return nil, nil
}
