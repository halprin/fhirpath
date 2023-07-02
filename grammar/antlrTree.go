package grammar

import "github.com/antlr4-go/antlr/v4"

type AntlrTree struct {
	text     string
	rule     string
	parent   *AntlrTree
	children []Tree
}

func NewAntlrTree(antlrTree antlr.RuleContext, parser *FhirpathParser) *AntlrTree {
	return newAntlrTreeWithParent(antlrTree, parser, nil)
}

func newAntlrTreeWithParent(antlrTree antlr.RuleContext, parser *FhirpathParser, parent *AntlrTree) *AntlrTree {
	tree := &AntlrTree{}

	tree.text = antlrTree.GetText()
	tree.rule = parser.GetRuleNames()[antlrTree.GetRuleIndex()]
	tree.parent = parent

	children := make([]Tree, 0, antlrTree.GetChildCount())

	for _, currentChild := range antlrTree.GetChildren() {
		payload, ok := currentChild.GetPayload().(antlr.RuleContext)

		if !ok {
			continue
		}

		child := newAntlrTreeWithParent(payload, parser, tree)

		children = append(children, child)
	}

	tree.children = children

	return tree
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
