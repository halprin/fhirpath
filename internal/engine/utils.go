package engine

import "reflect"

// convertInterfaceSliceToFhirOptionSlice converts a generic `[]interface{}` value to a slice of a FHIR option (`map[string]interface{}`)
// This is needed for some of the type casting in the execution engine.
func convertInterfaceSliceToFhirOptionSlice(interfaceSlice []interface{}) ([]map[string]interface{}, bool) {
	fhirOptions := make([]map[string]interface{}, 0, len(interfaceSlice))

	for _, interfaceValue := range interfaceSlice {
		fhirOption, ok := interfaceValue.(map[string]interface{})
		if !ok {
			return nil, ok
		}

		fhirOptions = append(fhirOptions, fhirOption)
	}

	return fhirOptions, true
}

// flatten flattens out any inner slices inside the passed in slice.
func flatten(slicesMaybe interface{}) []interface{} {
	var flattened []interface{}

	reflectedSlicesMaybe := reflect.ValueOf(slicesMaybe)

	for sliceIndex := 0; sliceIndex < reflectedSlicesMaybe.Len(); sliceIndex++ {
		currentValue := reflectedSlicesMaybe.Index(sliceIndex).Interface()

		if reflect.TypeOf(currentValue).Kind() == reflect.Slice {
			flattened = append(flattened, flatten(currentValue)...)
		} else {
			flattened = append(flattened, currentValue)
		}
	}

	return flattened
}
