package fhirpath

import (
	"encoding/json"
)

func Evaluate(fhir string, fhirPath string) ([]string, error) {
	
	_, err := unmarshalFhir(fhir)
	if err != nil {
		return nil, err
	}
	
	return []string{""}, nil
}

func unmarshalFhir(fhir string) (map[string]interface{}, error) {
	var fhirObject map[string]interface{}

	err := json.Unmarshal([]byte(fhir), &fhirObject)
	if err != nil {
		return nil, err
	}

	return fhirObject, nil
}
