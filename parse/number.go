package parse

import (
	"fmt"
	"github.com/halprin/fhirpath/lex"
	"strconv"
)

type Number struct {
	ValueInt   int
	ValueFloat float64
}

func NewNumber(buffer TokenBuffer) (Number, error) {
	integerToken, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Number{}, err
	}

	if integerToken.Type != lex.NUMBER {
		buffer.Push()
		return Number{}, NoGrammarParse
	}

	periodToken, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Number{}, err
	}

	if periodToken.Type != lex.PERIOD {
		//we're done; the number is an integer, not a decimal
		buffer.Push()
		intValue, err := strconv.Atoi(integerToken.Literal)
		if err != nil {
			return Number{}, fmt.Errorf("tried to make an integer even though lexxer found a number: %w", err)
		}
		return Number{ValueInt: intValue}, nil
	}

	decimalToken, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Number{}, err
	}

	if decimalToken.Type != lex.NUMBER {
		buffer.Push()
		return Number{}, NoGrammarParse
	}

	floatValue, err := strconv.ParseFloat(integerToken.Literal+periodToken.Literal+decimalToken.Literal, 64)
	if err != nil {
		return Number{}, fmt.Errorf("tried to make a float even though lexxer found a number: %w", err)
	}
	return Number{ValueFloat: floatValue}, nil
}
