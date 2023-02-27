package parse

import (
	"github.com/halprin/fhirpath/lex"
	"io"
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

func (receiver *TokenBuffer) PushTimes(times int) {
	receiver.bufferIndex -= times
}

func (receiver *TokenBuffer) PopUntilNot(tokenType int) ([]lex.Token, error) {
	var accumulator []lex.Token
	var token lex.Token
	var err error

	for {
		token, err = receiver.Pop()
		if err != nil {
			break
		} else if token.Type != tokenType {
			receiver.Push() //push that token back since it wasn't want we're looking for
			break
		}

		accumulator = append(accumulator, token)
	}

	if err == io.EOF && len(accumulator) > 0 {
		//we encountered the end, but we've read something means we have something valid; hide the error because it will be encountered in the next read
		err = nil
	}

	return accumulator, err
}
