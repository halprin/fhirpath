package grammar

type Tree interface {
	Text() string
	Rule() string
	Parent() Tree
	Children() []Tree
}
