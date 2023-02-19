package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
)

type HexDigit struct {
	Value rune
}

func NewHexDigit(buffer parse.TokenBuffer) (HexDigit, error) {
	token, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return HexDigit{}, err
	}

	if token.Type != lex.NUMBER && token.Type != lex.ALPHA_NUMERIC {
		buffer.Push()
		return HexDigit{}, parse.NoGrammarParse
	}

	if token.Type == lex.ALPHA_NUMERIC && !((token.Literal[0] >= 'a' && token.Literal[0] <= 'f') || (token.Literal[0] >= 'A' && token.Literal[0] <= 'F')) {
		//the character isn't between 0-9, a-f, or A-F
		buffer.Push()
		return HexDigit{}, parse.NoGrammarParse
	}

	if len(token.Literal) > 1 {
		//we need to split this token since it is longer than just one character; push the split back onto the buffer
		buffer.PushToken(lex.Token{
			Type:    token.Type,
			Literal: token.Literal[1:len(token.Literal)],
		})
	}

	return HexDigit{Value: rune(token.Literal[0])}, nil
}
