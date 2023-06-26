package listener

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
)

type FhirPathListener struct {
	*grammar.BaseFhirpathListener
}

func NewFhirPathListener() *FhirPathListener {
	return new(FhirPathListener)
}

func (listener *FhirPathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("EnterEveryRule: %d=%s\n", ctx.GetRuleIndex(), ctx.GetText())
}

func (listener *FhirPathListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("VisitTerminal: %s\n", node.GetText())
}
