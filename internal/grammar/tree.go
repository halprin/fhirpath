// Package grammar is used to create a `Tree` structure given a FHIR path to be parsed.
package grammar

// Tree represents a parsed FHIR path.  It is library agnostic.
type Tree interface {
	// Text contains the FHIR path represented by the specific grammar rule.
	Text() string
	// TerminalTexts returns a slice of strings representing tokens that terminate at this node.
	TerminalTexts() []string
	// Rule is the name of the grammar rule.
	Rule() string
	// Parent points to the parent node in the tree structure.
	Parent() Tree
	// Children point to this node's children trees.
	Children() []Tree
}
