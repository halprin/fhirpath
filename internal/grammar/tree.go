package grammar

type Tree interface {
	Text() string
	TerminalTexts() []string
	Rule() string
	Parent() Tree
	Children() []Tree
}
