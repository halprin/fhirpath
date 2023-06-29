package engine

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
)

type FhirPathEngine struct {
	*grammar.BaseFhirpathListener

	parser *grammar.FhirpathParser
	fhir   map[string]interface{}
}

func NewFhirPathEngine(fhir map[string]interface{}, parser *grammar.FhirpathParser) *FhirPathEngine {
	return &FhirPathEngine{
		parser: parser,
		fhir: fhir,
	}
}

func (engine *FhirPathEngine) Result() ([]interface{}, error) {
	return nil, nil
}

func (engine *FhirPathEngine) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("EnterEveryRule: %s: %s\n", engine.parser.GetRuleNames()[ctx.GetRuleIndex()], ctx.GetText())
	ctx.GetPayload()
}

func (engine *FhirPathEngine) VisitTerminal(node antlr.TerminalNode) {
//	fmt.Printf("VisitTerminal: %s, %s\n", node.GetText(), node.ToStringTree(nil, nil))
}
