package context

import (
	_ "embed"
	"encoding/json"
)

//go:generate go run generate.go

type Definition struct {
	Version   string                   `json:"version"`
	Resources []ResourceTypeDefinition `json:"resources"`
}

type ResourceTypeDefinition struct {
	Name   string `json:"name"`
	Fields []FieldTypes
}

type FieldTypes struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

//go:embed R5.json
var r5json string

var r5 Definition

func init() {
	var definition Definition
	err := json.Unmarshal([]byte(r5json), &definition)
	if err != nil {
		panic(err)
	}
}

func R5() Definition {
	return r5
}
