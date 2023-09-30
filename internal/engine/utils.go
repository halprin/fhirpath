package engine

// convertInterfaceSliceToFhirOptionSlice converts a generic `[]interface{}` value to a slice of a FHIR option (`map[string]interface{}`)
// This is needed for some of the type casting in the execution engine.
func convertInterfaceSliceToFhirOptionSlice(interfaceSlice []interface{}) []map[string]interface{} {
	fhirOptions := make([]map[string]interface{}, 0, len(interfaceSlice))

	for _, interfaceValue := range interfaceSlice {
		fhirOption := interfaceValue.(map[string]interface{})
		fhirOptions = append(fhirOptions, fhirOption)
	}

	return fhirOptions
}

// flatten flattens out any inner slices inside the passed in slice.
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
