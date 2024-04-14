package engine

import (
	"github.com/halprin/fhirpath/context"
	"github.com/halprin/fhirpath/internal/grammar"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
	"unicode/utf8"
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

	return NewDynamicValue(filterAndMap(fhirOptions, identifier, context)), nil
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

func filterAndMap(fhirOptions []map[string]interface{}, identifier string, context context.Definition) []interface{} {
	var filteredOptions []interface{}

	for _, currentFhirOption := range fhirOptions {
		value, ok := currentFhirOption[identifier]
		if ok {
			filteredOptions = append(filteredOptions, value)
			continue
		}

		polymorphismValue := polymorphismSearch(currentFhirOption, context, identifier)
		filteredOptions = append(filteredOptions, polymorphismValue)
	}

	//the filtered options could contain a slice itself, so those need to be unwrapped
	return flatten(filteredOptions)
}

func polymorphismSearch(currentFhirOption map[string]interface{}, context context.Definition, identifier string) interface{} {
	for _, resource := range context.Resources {
		for _, field := range resource.Fields {
			endOfFieldName := field.Name[len(field.Name)-3:]
			if endOfFieldName != "[x]" {
				//this isn't a field that has multiple types, so we don't need to test it
				continue
			}

			fieldNameWithoutPrefix := field.Name[strings.LastIndex(field.Name, ".")+1 : len(field.Name)-3]
			if fieldNameWithoutPrefix == identifier {
				for _, aType := range field.Types {
					titler := cases.Title(language.English)
					aType = titler.String(aType)
					value, ok := currentFhirOption[identifier+aType]
					if ok {
						return value
					}
				}
			}
		}
	}
	return ""
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
