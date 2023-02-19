package linkedList

import "github.com/halprin/fhirpath/lex"

type Node struct {
	Value    lex.Token
	Previous *Node
	Next     *Node
}

func (receiver *Node) InsertAfter(token lex.Token) *Node {
	nextNode := &Node{
		Value:    token,
		Previous: receiver,
		Next:     receiver.Next,
	}

	if receiver.Next != nil {
		receiver.Next.Previous = nextNode
	}

	receiver.Next = nextNode

	return nextNode
}
