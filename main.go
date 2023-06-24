package fhirpath

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/halprin/fhirpath/antlrgen"
)

func Stuff() {
	input := antlr.NewInputStream("Bundle.stuff")
	lexer := antlrgen.NewfhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := antlrgen.NewfhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(New)
}
