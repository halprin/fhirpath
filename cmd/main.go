package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/antlrgen"
)

type TreeShapeListener struct {
	*antlrgen.BaseFhirpathListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (listener *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewInputStream("Bundle.stuff")
	lexer := antlrgen.NewFhirpathLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := antlrgen.NewFhirpathParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Expression()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
