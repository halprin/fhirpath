package gammar

import (
	"fmt"
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"strconv"
)

type Number struct {
	ValueInt   int
	ValueFloat float64
}

func NewNumber(buffer parse.TokenBuffer) (Number, error) {
	integerTokens, err := buffer.PopUntilNot(lex.NUMERIC)
	if err != nil {
		buffer.PushTimes(len(integerTokens))
		return Number{}, err
	}

	if len(integerTokens) == 0 {
		return Number{}, parse.NoGrammarParse
	}

	periodToken, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return Number{}, err
	}

	if periodToken.Type != lex.PERIOD {
		//we're done; the number is an integer, not a decimal
		buffer.Push()
		intValue, err := strconv.Atoi(concatTokenLiterals(integerTokens))
		if err != nil {
			return Number{}, fmt.Errorf("tried to make an integer even though lexxer found a number: %w", err)
		}
		return Number{ValueInt: intValue}, nil
	}

	decimalTokens, err := buffer.PopUntilNot(lex.NUMERIC)
	if err != nil {
		buffer.PushTimes(len(decimalTokens))
		return Number{}, err
	}

	if len(decimalTokens) == 0 {
		return Number{}, parse.NoGrammarParse
	}

	floatValue, err := strconv.ParseFloat(concatTokenLiterals(integerTokens)+periodToken.Literal+concatTokenLiterals(decimalTokens), 64)
	if err != nil {
		return Number{}, fmt.Errorf("tried to make a float even though lexxer found a number: %w", err)
	}
	return Number{ValueFloat: floatValue}, nil
}
