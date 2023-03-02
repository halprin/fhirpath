package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"io"
)

type Escape struct {
	Character rune
}

func NewEscape(buffer *parse.TokenBuffer) (Escape, error) {
	backSlash, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Escape{}, err
	}

	if backSlash.Type != lex.BACK_SLASH {
		buffer.Push()
		return Escape{}, parse.NoGrammarParse
	}

	unicode, err := NewUnicode(buffer)
	if err == nil {
		return Escape{Character: unicode.Character}, nil
	} else if err == io.EOF {
		buffer.Push()
		return Escape{}, err
	}

	escapedCharacter, err := buffer.Pop()
	if err != nil {
		buffer.PushTimes(2)
		return Escape{}, err
	}

	if escapedCharacter.Type == lex.BACK_TICK {
		return Escape{Character: escapedCharacter.Literal}, nil
	} else if escapedCharacter.Type == lex.QUOTE {
		return Escape{Character: escapedCharacter.Literal}, nil
	} else if escapedCharacter.Type == lex.BACK_SLASH {
		return Escape{Character: escapedCharacter.Literal}, nil
	} else if escapedCharacter.Type == lex.QUOTE {
		return Escape{Character: escapedCharacter.Literal}, nil
	} else if escapedCharacter.Type == lex.ALPHA {
		if escapedCharacter.Literal == 'f' {
			return Escape{Character: '\f'}, nil
		} else if escapedCharacter.Literal == 'n' {
			return Escape{Character: '\n'}, nil
		} else if escapedCharacter.Literal == 'r' {
			return Escape{Character: '\r'}, nil
		} else if escapedCharacter.Literal == 't' {
			return Escape{Character: '\t'}, nil
		}
	}

	buffer.PushTimes(2)
	return Escape{}, parse.NoGrammarParse
}
