package engine

import (
	"github.com/halprin/fhirpath/grammar"
	"github.com/halprin/rangechain"
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
		filteredFhirOptionsInterface, err := rangechain.FromSlice(fhirOptions).Filter(func(currentFhirOptionInterface interface{}) (bool, error) {
			currentFhirOption := currentFhirOptionInterface.(map[string]interface{})

			resourceTypeInterface, ok := currentFhirOption["resourceType"]
			if !ok {
				return false, nil
			}

			resourceType := resourceTypeInterface.(string)

			return resourceType == identifier, nil
		}).Slice()

		if err != nil {
			return nil, err
		}

		return convertInterfaceSliceToFhirOptionSlice(filteredFhirOptionsInterface), nil
	}

	//a filter on some generic field.  I.e. gender.
	return rangechain.FromSlice(fhirOptions).Filter(func(currentFhirOptionInterface interface{}) (bool, error) {
		currentFhirOption := currentFhirOptionInterface.(map[string]interface{})

		_, ok := currentFhirOption[identifier]
		return ok, nil
	}).Map(func(currentFhirOptionInterface interface{}) (interface{}, error) {
		currentFhirOption := currentFhirOptionInterface.(map[string]interface{})

		fieldValueInterface := currentFhirOption[identifier]

		return fieldValueInterface, nil
	}).Slice()
}

func convertInterfaceSliceToFhirOptionSlice(interfaceSlice []interface{}) []map[string]interface{} {
	fhirOptions := make([]map[string]interface{}, 0, len(interfaceSlice))
	
	for _, interfaceValue := range interfaceSlice {
		fhirOption := interfaceValue.(map[string]interface{})
		fhirOptions = append(fhirOptions, fhirOption)
	}

	return fhirOptions
}
