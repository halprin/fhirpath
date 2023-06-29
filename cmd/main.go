package main

import (
	_ "embed"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
	"github.com/halprin/fhirpath/listener"
)

func main() {
	input := antlr.NewInputStream("name.where(use='usual' or use='official').given.first()")
	lexer := grammar.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := grammar.NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := parser.Expression()
	antlr.ParseTreeWalkerDefault.Walk(listener.NewFhirPathListener("patientFhir", parser), tree)
}
