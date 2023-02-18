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
