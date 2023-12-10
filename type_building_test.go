package fhirpath

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type structureDefinition struct {
	Name     string `json:"name"`
	Snapshot struct {
		Element []struct {
			Path string `json:"path"`
			Type []struct {
				Code string `json:"code"`
			} `json:"type"`
		} `json:"element"`
	} `json:"snapshot"`
}

func TestTypeBuilding(t *testing.T) {

	fhirVersions := []string{
		"R5",
		"R4B",
		"R4",
		"STU3",
		"DSTU2",
	}

	for _, version := range fhirVersions {
		err := constructTypesForFhirVersion(version)
		if err != nil {
			t.Error(err)
		}
	}
}

func constructTypesForFhirVersion(fhirVersion string) error {
	resources := []string{
		"Patient",
		"Encounter",
		"Observation",
		// Add more resource names here
	}

	for _, resource := range resources {
		resourceDefinition, err := fetchStructureDefinition(fhirVersion, resource)
		if err != nil {
			return err
		}

		for _, elem := range resourceDefinition.Snapshot.Element {
			fmt.Printf("%s %s types =\n", fhirVersion, elem.Path)
			for _, currentType := range elem.Type {
				fmt.Printf("\t%s\n", currentType.Code)
			}
		}
	}

	return nil
}

func fetchStructureDefinition(fhirVersion string, resource string) (structureDefinition, error) {
	url := fmt.Sprintf("https://hl7.org/fhir/%s/%s.profile.json", fhirVersion, resource)

	response, err := http.Get(url)
	if err != nil {
		return structureDefinition{}, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return structureDefinition{}, err
	}

	var resourceDefinition structureDefinition
	err = json.Unmarshal(responseBody, &resourceDefinition)
	if err != nil {
		return structureDefinition{}, err
	}

	return resourceDefinition, nil
}
