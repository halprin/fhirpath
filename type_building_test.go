package fhirpath

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
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
		//"R4B",
		//"R4",
		//"STU3",
		//"DSTU2",
	}

	for _, version := range fhirVersions {
		err := constructTypesForFhirVersion(version)
		if err != nil {
			t.Error(err)
		}
	}
}

func constructTypesForFhirVersion(fhirVersion string) error {
	log.Printf("Constructing types for FHIR version %s\n", fhirVersion)

	resources, err := getResourcesForVersion(fhirVersion)
	if err != nil {
		return err
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

func getResourcesForVersion(fhirVersion string) ([]string, error) {
	log.Printf("Downloading FHIR resource list for version %s\n", fhirVersion)

	url := fmt.Sprintf("https://hl7.org/fhir/%s/resourcelist.html", fhirVersion)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resources := make([]string, 0)

	log.Printf("Parsing FHIR resources for FHIR version %s\n", fhirVersion)

	tokenizer := html.NewTokenizer(response.Body)

	var isInAlphabeticalDivTag bool
	var isInResourceATag bool

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken:
			return resources, nil
		case tokenType == html.StartTagToken:
			token := tokenizer.Token()
			if isInAlphabeticalDivTag && token.Data == "a" {
				for _, attribute := range token.Attr {
					if attribute.Key == "title" && attribute.Val != "Maturity Level" && attribute.Val != "Normative Content" {
						isInResourceATag = true
					}
				}
			}
			if token.Data == "div" {
				for _, attribute := range token.Attr {
					if attribute.Key == "id" && attribute.Val == "tabs-2" { // tab-2 is the alphabetical listing of resources
						isInAlphabeticalDivTag = true
					}
				}
			}
		case tokenType == html.EndTagToken:
			token := tokenizer.Token()
			if isInAlphabeticalDivTag && token.Data == "div" {
				isInAlphabeticalDivTag = false
			}
			if isInResourceATag && token.Data == "a" {
				isInResourceATag = false
			}
		case tokenType == html.TextToken:
			if isInResourceATag {
				token := tokenizer.Token()
				resources = append(resources, token.Data)
			}
		}
	}
}

func fetchStructureDefinition(fhirVersion string, resource string) (structureDefinition, error) {
	log.Printf("Fetching structure definition for FHIR version %s, resource %s\n", fhirVersion, resource)

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

	log.Printf("Parsing structure definition for FHIR version %s, resource %s\n", fhirVersion, resource)

	var resourceDefinition structureDefinition
	err = json.Unmarshal(responseBody, &resourceDefinition)
	if err != nil {
		return structureDefinition{}, err
	}

	return resourceDefinition, nil
}
