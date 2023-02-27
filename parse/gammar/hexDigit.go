package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
)

type HexDigit struct {
	Digit rune
}

func NewHexDigit(buffer *parse.TokenBuffer) (HexDigit, error) {
	token, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return HexDigit{}, err
	}

	if token.Type != lex.NUMERIC && token.Type != lex.ALPHA {
		buffer.Push()
		return HexDigit{}, parse.NoGrammarParse
	}

	if token.Type == lex.ALPHA && (token.Literal >= 'g' && token.Literal <= 'z') || (token.Literal >= 'G' && token.Literal <= 'Z') {
		buffer.Push()
		return HexDigit{}, parse.NoGrammarParse
	}

	return HexDigit{Digit: token.Literal}, nil
}
