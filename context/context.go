package context

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

//go:

//func R5() Definition {
//
//}
