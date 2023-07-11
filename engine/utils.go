package engine

//used to convert a generic `[]interface{}` value to a slice of a FHIR option (`map[string]interface{}`)
//this is needed for some of the type casting in the execution engine.  E.g. `InvocationExpression`.
func convertInterfaceSliceToFhirOptionSlice(interfaceSlice []interface{}) []map[string]interface{} {
	fhirOptions := make([]map[string]interface{}, 0, len(interfaceSlice))

	for _, interfaceValue := range interfaceSlice {
		fhirOption := interfaceValue.(map[string]interface{})
		fhirOptions = append(fhirOptions, fhirOption)
	}

	return fhirOptions
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
