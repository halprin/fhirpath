package grammar

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/internal/grammar/antlrGen"
)

// CreateTree is the top-level function to parse a FHIR path into a `Tree`.
func CreateTree(fhirPath string) (Tree, error) {
	//in this specific case, the ANTLR library is used to parse the FHIR path
	input := antlr.NewInputStream(fhirPath)
	lexer := antlrGen.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := antlrGen.NewFhirpathParser(stream)
	errorExtractor := &AntlrErrorHanderExtractor{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorExtractor)
	tree := parser.Expression()

	return NewAntlrTree(tree), errorExtractor.Error()
}
