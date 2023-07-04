package fhirpath

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed sample/patient.json
var fhirPatient string

func TestEvaluate(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.gender")
	
	assert.NoError(t, err)
	assert.Contains(t, result, "female")
}
