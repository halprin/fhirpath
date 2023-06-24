package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
	"github.com/halprin/fhirpath/listener"
)

func main() {
	input := antlr.NewInputStream("Bundle.stuff")
	lexer := grammar.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Expression()
	antlr.ParseTreeWalkerDefault.Walk(listener.NewFhirPathListener(), tree)
}
