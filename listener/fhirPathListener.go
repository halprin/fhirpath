package listener

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
)

type FhirPathListener struct {
	*grammar.BaseFhirpathListener

	parser *grammar.FhirpathParser
}

func NewFhirPathListener(parser *grammar.FhirpathParser) *FhirPathListener {
	return &FhirPathListener{parser: parser}
}

func (listener *FhirPathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("EnterEveryRule: %s: %s\n", listener.parser.GetRuleNames()[ctx.GetRuleIndex()], ctx.GetText())
	ctx.GetPayload()
}

func (listener *FhirPathListener) VisitTerminal(node antlr.TerminalNode) {
//	fmt.Printf("VisitTerminal: %s, %s\n", node.GetText(), node.ToStringTree(nil, nil))
}
