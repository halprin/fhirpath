package grammar

import (
	"github.com/antlr4-go/antlr/v4"
)

//go:generate ./generate.sh

func CreateTree(fhirPath string) (Tree, error) {
	input := antlr.NewInputStream(fhirPath)
	lexer := NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := parser.Expression()

	return NewAntlrTree(tree), nil
}
