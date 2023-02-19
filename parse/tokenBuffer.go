package parse

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/linkedList"
)

type TokenBuffer struct {
	Lexer  lex.Lexer
	buffer *linkedList.Node
}

func NewTokenBuffer(lexer lex.Lexer) TokenBuffer {
	return TokenBuffer{
		Lexer:  lexer,
		buffer: &linkedList.Node{},
	}
}

func (receiver *TokenBuffer) Pop() (lex.Token, error) {

	if receiver.buffer.Next != nil {
		//we've backtracked and need to return a buffered token
		receiver.buffer = receiver.buffer.Next
		return receiver.buffer.Value, nil
	}

	token, err := receiver.Lexer.NextToken()
	if err != nil {
		return lex.Token{}, err
	}

	receiver.buffer = receiver.buffer.InsertAfter(token)

	return token, nil
}

func (receiver *TokenBuffer) Push() {
	if receiver.buffer.Previous == nil {
		return
	}

	receiver.buffer = receiver.buffer.Previous
}

func (receiver *TokenBuffer) PushToken(token lex.Token) {
	receiver.buffer.InsertAfter(token)
}
