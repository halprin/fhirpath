package fhirpath

import (
	_ "embed"
	"reflect"
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

func TestEvaluate_ParseError(t *testing.T) {
	_, err := Evaluate[string](fhirPatient, "Patient..gender")

	assert.Error(t, err)
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

func TestEvaluate_Exists(t *testing.T) {
	result, err := Evaluate[bool](fhirPatient, "Patient.name.exists()")

	assert.NoError(t, err)
	assert.Contains(t, result, true)
}

func TestEvaluate_Exists_False(t *testing.T) {
	result, err := Evaluate[bool](fhirPatient, "Patient.deceased.exists()")

	assert.NoError(t, err)
	assert.Contains(t, result, false)
}

func TestEvaluate_Exists_InnerWhere(t *testing.T) {
	result, err := Evaluate[bool](fhirPatient, "Patient.name.exists(use = 'official')")

	assert.NoError(t, err)
	assert.Contains(t, result, true)
}

func TestEvaluate_Equality_Boolean(t *testing.T) {
	result, err := Evaluate[bool](fhirPatient, "Patient.name.exists() = true")

	assert.NoError(t, err)
	assert.Contains(t, result, true)
}

func TestEvaluate_Polymorphism_Exact(t *testing.T) {
	//TODO: this test fails because we construct an `int` for the `1`, but Go unmarshals the `1` in
	//TODO: `multipleBirthInteger` as a `float64`.  The answer is probably to update `NumberLiteral` to always construct
	//TODO: a `float64` regardless of whether there is a `.`.
	result, err := Evaluate[bool](fhirPatient, "Patient.multipleBirthInteger = 1")

	assert.NoError(t, err)
	assert.Contains(t, result, true)
}

func TestEvaluate_Polymorphism_Inprecise(t *testing.T) {
	result, err := Evaluate[bool](fhirPatient, "Patient.multipleBirth = 1")

	assert.NoError(t, err)
	assert.Contains(t, result, true)
}

func TestOfficial_testSliceOfBool(t *testing.T) {
	booleans := []bool{true, false, true}
	var obfuscatedBooleans interface{}
	obfuscatedBooleans = booleans

	reflectedSlice := reflect.ValueOf(obfuscatedBooleans)

	for i := 0; i < reflectedSlice.Len(); i++ {
		t.Log(reflectedSlice.Index(i).Interface())
	}

}
