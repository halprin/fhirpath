package grammar

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/internal/grammar/antlrGen"
	"reflect"
	"strings"
)

type AntlrTree struct {
	text          string
	terminalTexts []string
	rule          string
	parent        *AntlrTree
	children      []Tree
}

func NewAntlrTree(antlrTree antlr.RuleContext) *AntlrTree {
	return newAntlrTreeWithParent(antlrTree, nil)
}

func newAntlrTreeWithParent(antlrTree antlr.RuleContext, parent *AntlrTree) *AntlrTree {
	tree := &AntlrTree{
		text: antlrTree.GetText(),
		rule: trimmedAntlrType(reflect.TypeOf(antlrTree).String()),
		parent: parent,
	}

	children := make([]Tree, 0, antlrTree.GetChildCount())
	var terminalTexts []string

	for _, currentChild := range antlrTree.GetChildren() {
		payload, ok := currentChild.(antlr.RuleContext)

		if !ok {
			//it could be a terminal node
			terminalPayload, ok := currentChild.(antlr.TerminalNode)
			if !ok {
				//the child was nothing we need to parse
				continue
			}

			terminalTexts = append(terminalTexts, terminalPayload.GetText())

			continue
		}

		child := newAntlrTreeWithParent(payload, tree)

		children = append(children, child)
	}

	tree.children = children
	tree.terminalTexts = terminalTexts

	return tree
}

func trimmedAntlrType(antlrType string) string {

	//get the package name used by the generated ANTLR code at run time
	antlrGenPackagePath := reflect.TypeOf(antlrGen.FhirpathParser{}).PkgPath()
	packageParts := strings.Split(antlrGenPackagePath, "/")
	antlrGenPackageName := packageParts[len(packageParts) - 1]

	startIndex := len("*" + antlrGenPackageName + ".")  //trim the beginning *packageName.
	endIndex := len("Context")  //trim the ending Context

	return antlrType[startIndex:len(antlrType) - endIndex]
}

func (receiver *AntlrTree) Text() string {
	return receiver.text
}

func (receiver *AntlrTree) TerminalTexts() []string {
	return receiver.terminalTexts
}

func (receiver *AntlrTree) Rule() string {
	return receiver.rule
}

func (receiver *AntlrTree) Children() []Tree {
	return receiver.children
}

func (receiver *AntlrTree) Parent() Tree {
	return receiver.parent
}
