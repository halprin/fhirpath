package grammar

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/internal/grammar/antlrGen"
)

func CreateTree(fhirPath string) (Tree, error) {
	input := antlr.NewInputStream(fhirPath)
	lexer := antlrGen.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := antlrGen.NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := parser.Expression()

	return NewAntlrTree(tree), nil
}
