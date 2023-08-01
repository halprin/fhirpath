package fhirpath

import (
	_ "embed"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed official_tests/r4/tests-fhir-r4.xml
var officialTestXmlSpec []byte

type OfficialTests struct {
	Groups []OfficialTestGroup
}

type OfficialTestGroup struct {
	Name  string
	Tests []OfficialTest
}

type OfficialTest struct {
	Name       string
	Expression string
	InputFile  string
	Outputs    []string
}

func TestOfficial(t *testing.T) {
	//parse official_tests/r4/test-fhir-r4.xml
	var officialTests OfficialTests
	err := xml.Unmarshal(officialTestXmlSpec, &officialTests)
	assert.NoError(t, err)
	//for each test, call officialTestTemplate(...) and then call t.Run() the return value
}

func officialTestTemplate(testName string, fhirPath string, fhir string, expectedResult []string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := Evaluate[any](fhir, fhirPath)
		if err != nil {
			t.Logf("Evaluate failed with an error: %s", err.Error())
		}
	}
}
