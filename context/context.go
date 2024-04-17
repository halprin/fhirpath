package context

import (
	_ "embed"
	"encoding/json"
)

//go:generate go run generate.go

type Definition struct {
	Version   string                            `json:"version"`
	Resources map[string]ResourceTypeDefinition `json:"resources"`
}

type ResourceTypeDefinition struct {
	Name   string `json:"name"`
	Fields map[string]FieldTypes
}

type FieldTypes struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

//go:embed R5.json
var r5json string

//go:embed R4B.json
var r4bjson string

//go:embed R4.json
var r4json string

//go:embed STU3.json
var stu3json string

//go:embed DSTU2.json
var dstu2json string

var r5 Definition
var r4b Definition
var r4 Definition
var stu3 Definition
var dstu2 Definition

func init() {

	if len(r5json) == 0 {
		return
	}

	err := json.Unmarshal([]byte(r5json), &r5)
	if err != nil {
		panic(err)
	}

	if len(r4bjson) == 0 {
		return
	}

	err = json.Unmarshal([]byte(r4bjson), &r4b)
	if err != nil {
		panic(err)
	}

	if len(r4json) == 0 {
		return
	}

	err = json.Unmarshal([]byte(r4json), &r4)
	if err != nil {
		panic(err)
	}

	if len(stu3json) == 0 {
		return
	}

	err = json.Unmarshal([]byte(stu3json), &stu3)
	if err != nil {
		panic(err)
	}

	if len(dstu2json) == 0 {
		return
	}

	err = json.Unmarshal([]byte(dstu2json), &dstu2)
	if err != nil {
		panic(err)
	}
}

func R5() Definition {
	return r5
}

func R4B() Definition {
	return r4b
}

func R4() Definition {
	return r4
}

func STU3() Definition {
	return stu3
}

func DSTU2() Definition {
	return dstu2
}
