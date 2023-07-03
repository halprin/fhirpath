package grammar

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/halprin/fhirpath/grammar/antlrGen"
	"reflect"
	"strings"
)

type AntlrTree struct {
	text     string
	rule     string
	parent   *AntlrTree
	children []Tree
}

func NewAntlrTree(antlrTree antlr.RuleContext) *AntlrTree {
	return newAntlrTreeWithParent(antlrTree, nil)
}

func newAntlrTreeWithParent(antlrTree antlr.RuleContext, parent *AntlrTree) *AntlrTree {
	tree := &AntlrTree{}

	tree.text = antlrTree.GetText()
	tree.rule = trimmedAntlrType(reflect.TypeOf(antlrTree).String())
	tree.parent = parent

	children := make([]Tree, 0, antlrTree.GetChildCount())

	for _, currentChild := range antlrTree.GetChildren() {
		payload, ok := currentChild.(antlr.RuleContext)

		if !ok {
			continue
		}

		child := newAntlrTreeWithParent(payload, tree)

		children = append(children, child)
	}

	tree.children = children

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

func (receiver *AntlrTree) Rule() string {
	return receiver.rule
}

func (receiver *AntlrTree) Children() []Tree {
	return receiver.children
}

func (receiver *AntlrTree) Parent() Tree {
	return receiver.parent
}
