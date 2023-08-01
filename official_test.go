package fhirpath

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed official_tests/r4/tests-fhir-r4.xml
var officialTestXmlSpec []byte

type OfficialTests struct {
	XMLName xml.Name            `xml:"tests"`
	Groups  []OfficialTestGroup `xml:"group"`
}

type OfficialTestGroup struct {
	Name  string         `xml:"name,attr"`
	Tests []OfficialTest `xml:"test"`
}

type OfficialTest struct {
	Name       string `xml:"name,attr"`
	Expression string `xml:"expression"`
	InputFile  string `xml:"inputfile,attr"`
	//	Invalid    bool     `xml:"expression>invalid,attr"`
	Outputs []string `xml:"output"`
}

func TestOfficial(t *testing.T) {
	//parse official_tests/r4/test-fhir-r4.xml
	var officialTests OfficialTests
	err := xml.Unmarshal(officialTestXmlSpec, &officialTests)
	assert.NoError(t, err)
	//for each test, call officialTestTemplate(...) and then call t.Run() the return value
	for _, group := range officialTests.Groups {
		for _, test := range group.Tests {
			testName := fmt.Sprintf("%s/%s", group.Name, test.Name)
			t.Run(testName, officialTestTemplate(test.Expression, "", test.Outputs))
		}
	}
}

func officialTestTemplate(fhirPath string, fhir string, expectedResult []string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := Evaluate[any](fhir, fhirPath)
		if err != nil {
			t.Logf("Evaluate failed with an error: %s", err.Error())
		}
	}
}
