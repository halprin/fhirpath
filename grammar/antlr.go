package grammar

import (
	"github.com/antlr4-go/antlr/v4"
)

//go:generate ./generate.sh

func AntlrExecute[T any](fhir map[string]interface{}, fhirPath string) ([]T, error) {
	input := antlr.NewInputStream(fhirPath)
	lexer := NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := parser.Expression()
	
	fhirEngine := NewFhirPathEngine[T](fhir, parser)
	
	antlr.ParseTreeWalkerDefault.Walk(fhirEngine, tree)
	
	return fhirEngine.Result()
}
