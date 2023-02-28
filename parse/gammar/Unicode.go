package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"strconv"
)

type Unicode struct {
	Character rune
}

func NewUnicode(buffer *parse.TokenBuffer) (Unicode, error) {
	uCharacter, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Unicode{}, err
	}

	if uCharacter.Type != lex.ALPHA || uCharacter.Literal != 'u' {
		buffer.Push()
		return Unicode{}, parse.NoGrammarParse
	}

	hexDigit1, err := NewHexDigit(buffer)
	if err != nil {
		buffer.Push() //undo the u character
		return Unicode{}, err
	}

	hexDigit2, err := NewHexDigit(buffer)
	if err != nil {
		buffer.PushTimes(2) //undo the u character and the first hex digit
		return Unicode{}, err
	}

	hexDigit3, err := NewHexDigit(buffer)
	if err != nil {
		buffer.PushTimes(3) //undo the u character and the first - second hex digit
		return Unicode{}, err
	}

	hexDigit4, err := NewHexDigit(buffer)
	if err != nil {
		buffer.PushTimes(4) //undo the u character and the first - third hex digit
		return Unicode{}, err
	}

	concatenatedHexDigits := string(hexDigit1.Digit) + string(hexDigit2.Digit) + string(hexDigit3.Digit) + string(hexDigit4.Digit)
	hexInteger, err := strconv.ParseInt(concatenatedHexDigits, 16, 32)
	if err != nil {
		buffer.PushTimes(5) //undo the u character and the first - fourth hex digit
		return Unicode{}, parse.NoGrammarParse
	}

	return Unicode{Character: rune(hexInteger)}, nil
}
