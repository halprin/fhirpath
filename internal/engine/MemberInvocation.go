package engine

import (
	"github.com/halprin/fhirpath/context"
	"unicode"
	"unicode/utf8"

	"github.com/halprin/fhirpath/internal/grammar"
)

// MemberInvocation represents filtering based on a segment after a period in the FHIR path.
func (receiver *engine) MemberInvocation(fhirOptions []map[string]interface{}, node grammar.Tree, context context.Definition) (*DynamicValue, error) {
	identifierDynamicValue, err := receiver.Execute(fhirOptions, node.Children()[0], context)
	if err != nil {
		return nil, err
	}

	identifier := identifierDynamicValue.Value.(string)

	firstCharacter, _ := utf8.DecodeRuneInString(identifier)
	if unicode.IsUpper(firstCharacter) {
		//wants to filter on a specific resource type.  I.e. 'Patient'.  This is specific to the "resourceType" field.
		return NewDynamicValue(filterResourceType(fhirOptions, identifier)), nil
	}

	//a filter on some generic field.  I.e. gender.

	return NewDynamicValue(filterAndMap(fhirOptions, identifier)), nil
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
