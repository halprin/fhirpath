package fhirpath

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"github.com/halprin/fhirpath/context"
	"os"
	"reflect"
	"strings"
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
	Name string `xml:"name,attr"`
	//	Expression string   `xml:"expression"`
	Expression OfficialExpression `xml:"expression"`
	InputFile  string             `xml:"inputfile,attr"`
	Outputs    []string           `xml:"output"`
}

type OfficialExpression struct {
	Expression string `xml:",chardata"`
	Invalid    string `xml:"invalid,attr"`
}

func TestOfficial(t *testing.T) {
	//parse official_tests/r4/test-fhir-r4.xml
	var officialTests OfficialTests
	err := xml.Unmarshal(officialTestXmlSpec, &officialTests)
	assert.NoError(t, err)

	failTests := false
	totalTests := 0
	passedTests := 0

	for _, group := range officialTests.Groups {
		for _, test := range group.Tests {
			testName := fmt.Sprintf("%s/%s", group.Name, test.Name)
			totalTests++
			fhir, err := readFhirTestFile(convertXmlFileNameToJsonFileName(test.InputFile))
			assert.NoError(t, err)
			passed := t.Run(testName, officialTestTemplate(test.Expression, fhir, test.Outputs, failTests))
			if passed {
				passedTests++
			}
		}
	}

	t.Logf("%d/%d tests pass", passedTests, totalTests)
}

func convertXmlFileNameToJsonFileName(fileName string) string {
	return strings.Replace(fileName, ".xml", ".json", 1)
}

func readFhirTestFile(fileName string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("official_tests/r4/input/%s", fileName))
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func officialTestTemplate(expression OfficialExpression, fhir string, expectedResult []string, fail bool) func(*testing.T) {
	return func(t *testing.T) {

		//report on any possible panics
		defer func() {
			r := recover()
			if r != nil {
				t.Logf("Evaluate failed with a panic: %v", r)
				if fail {
					t.Fail()
				}
			}
		}()

		results, err := EvaluateWithContext[any](fhir, expression.Expression, context.R4())
		if err != nil {
			if expression.Invalid != "" {
				//this was an expected error
				t.Log("Successfully failed")
				return
			}
			t.Logf("Evaluate failed with an error: %s", err.Error())
			if fail {
				t.Fail()
			}
			return
		}

		stringifiedResults := stringifySlice(results)

		if len(expectedResult) != len(stringifiedResults) {
			t.Log("Expected results are not equal to actual results")
			t.Logf("Expected=%v", expectedResult)
			t.Logf("Actual=%v", stringifiedResults)
			if fail {
				t.Fail()
			}
			return
		}

		expectedCount := make(map[string]int)
		for _, currentExpectedResult := range expectedResult {
			count, ok := expectedCount[currentExpectedResult]
			if !ok {
				expectedCount[currentExpectedResult] = 1
				continue
			}

			expectedCount[currentExpectedResult] = count + 1
		}

		actualCount := make(map[string]int)
		for _, currentStringifiedResults := range stringifiedResults {
			count, ok := actualCount[currentStringifiedResults]
			if !ok {
				actualCount[currentStringifiedResults] = 1
				continue
			}

			actualCount[currentStringifiedResults] = count + 1
		}

		if !reflect.DeepEqual(expectedCount, actualCount) {
			t.Log("Expected results are not equal to actual results")
			t.Logf("Expected=%v", expectedResult)
			t.Logf("Actual=%v", stringifiedResults)
			if fail {
				t.Fail()
			}
		}
	}
}

func stringifySlice[T any](results []T) []string {
	stringValues := make([]string, len(results))

	for index, result := range results {
		stringValues[index] = fmt.Sprintf("%v", result)
	}

	return stringValues
}
