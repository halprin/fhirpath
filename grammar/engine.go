package grammar

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type FhirPathEngine[T any] struct {
	*BaseFhirpathListener

	parser         *FhirpathParser
	fhir           map[string]interface{}
	currentContext map[string]interface{}
}

func NewFhirPathEngine[T any](fhir map[string]interface{}, parser *FhirpathParser) *FhirPathEngine[T] {
	return &FhirPathEngine[T]{
		parser: parser,
		fhir:   fhir,
		currentContext: fhir,
	}
}

func (engine *FhirPathEngine[T]) Result() ([]T, error) {
	return nil, nil
}

func (engine *FhirPathEngine[T]) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("EnterEveryRule: %s: %s\n", engine.parser.GetRuleNames()[ctx.GetRuleIndex()], ctx.GetText())
}

func (engine *FhirPathEngine[T]) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("VisitTerminal: %s\n", node.GetText())
}
