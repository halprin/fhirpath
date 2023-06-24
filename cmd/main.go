package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar"
)

type TreeShapeListener struct {
	*grammar.BaseFhirpathListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewInputStream("Bundle.stuff")
	lexer := grammar.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Expression()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
