package fhirpath

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample/patient.json
var fhirPatient string

func TestEvaluate_Value(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.gender")

	assert.NoError(t, err)
	assert.Contains(t, result, "female")
}

func TestEvaluate_Where_Equal(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.identifier.where(system='http://new-republic.gov/galactic-citizen-identifier').value")

	assert.NoError(t, err)
	assert.Contains(t, result, "b531d827-de9a-4e2e-a53b-8621bd29f656")
}

func TestEvaluate_Where_NotEqual(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.identifier.where(system!='http://new-republic.gov/galactic-citizen-identifier').value")

	assert.NoError(t, err)
	assert.Contains(t, result, "S99955754")
}

func TestEvaluate_Where_Other(t *testing.T) {
	result, err := Evaluate[string](fhirBundleOrder, "Bundle.entry.where(resource.resourceType='ServiceRequest').resource.code.coding.code")

	assert.NoError(t, err)
	assert.Contains(t, result, "54089-8")
	assert.Contains(t, result, "57717-1")
}

func TestEvaluate_Index(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.identifier[1].value")

	assert.NoError(t, err)
	assert.Contains(t, result, "b531d827-de9a-4e2e-a53b-8621bd29f656")
}

func TestEvaluate_Index_NotInteger(t *testing.T) {
	_, err := Evaluate[string](fhirPatient, "Patient.identifier[2.6].value")

	assert.Error(t, err)
}

func TestEvaluate_Index_AboveSizeBecomesEmptyResult(t *testing.T) {
	result, err := Evaluate[string](fhirPatient, "Patient.identifier[2].value")

	assert.NoError(t, err)
	assert.Len(t, result, 0)
}
