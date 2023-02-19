package parse

import (
	"github.com/halprin/fhirpath/lex"
)

type TokenBuffer struct {
	Lexer       lex.Lexer
	buffer      []lex.Token
	bufferIndex int
}

func (receiver *TokenBuffer) Pop() (lex.Token, error) {
	if receiver.bufferIndex <= len(receiver.buffer)-1 {
		token := receiver.buffer[receiver.bufferIndex]
		receiver.bufferIndex++
		return token, nil
	}

	token, err := receiver.Lexer.NextToken()
	if err != nil {
		return lex.Token{}, err
	}

	receiver.buffer = append(receiver.buffer, token)
	receiver.bufferIndex++

	return token, nil
}

func (receiver *TokenBuffer) Push() {
	receiver.bufferIndex--
}

func (receiver *TokenBuffer) PushToken(token lex.Token) {
	//TODO: this won't work; we need a linked list or something because what if we insert when the bufferIndex points to the middle of the buffer?  We don't want to shift everything.
	receiver.buffer = append(receiver.buffer, token)
	receiver.bufferIndex--
}
