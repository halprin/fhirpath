package engine

import (
	"github.com/halprin/fhirpath/grammar"
	"unicode"
	"unicode/utf8"
)

func (receiver *engine) MemberInvocation(fhirOptions []map[string]interface{}, node grammar.Tree) (interface{}, error) {
	identifierInterface, err := receiver.Execute(fhirOptions, node.Children()[0])
	if err != nil {
		return nil, err
	}

	identifier := identifierInterface.(string)

	firstCharacter, _ := utf8.DecodeRuneInString(identifier)
	if unicode.IsUpper(firstCharacter) {
		//wants to filter on a specific resource type.  I.e. 'Patient'.  This is specific to the "resourceType" field.
		return filterResourceType(fhirOptions, identifier), nil
	}

	//a filter on some generic field.  I.e. gender.

	return filterAndMap(fhirOptions, identifier), nil
}

func filterResourceType(fhirOptions []map[string]interface{}, identifier string) []map[string]interface{} {
	var filteredFhirOptions []map[string]interface{}

	for _, currentFhirOption := range fhirOptions {
		if fhirOptionHasRequestedFieldValue(currentFhirOption, "resourceType", identifier) {
			filteredFhirOptions = append(filteredFhirOptions, currentFhirOption)
		}
	}

	return filteredFhirOptions
}

func filterAndMap(fhirOptions []map[string]interface{}, identifier string) []interface{} {
	var filteredOptions []interface{}

	for _, currentFhirOption := range fhirOptions {
		value, ok := currentFhirOption[identifier]
		if ok {
			filteredOptions = append(filteredOptions, value)
		}
	}

	//the filtered options could contain a slice itself, so those need to be unwrapped
	return flatten(filteredOptions)
}

func fhirOptionHasRequestedFieldValue[T comparable](fhirOption map[string]interface{}, fieldName string, fieldValue T) bool {
	fhirValueInterface, ok := fhirOption[fieldName]
	if !ok {
		return false
	}

	fhirValue, ok := fhirValueInterface.(T)
	if !ok {
		return false
	}

	return fhirValue == fieldValue
}

func flatten(slices []interface{}) []interface{} {
	var flattened []interface{}

	for _, currentPossibleSlice := range slices {

		currentSlice, ok := currentPossibleSlice.([]interface{})

		if !ok {
			flattened = append(flattened, currentPossibleSlice)
			continue
		}

		flattened = append(flattened, currentSlice...)
	}

	return flattened
}
