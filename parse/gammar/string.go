package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"strings"
)

type String struct {
	Value string
}

func NewString(buffer *parse.TokenBuffer) (String, error) {
	startQuoteToken, err := buffer.Pop()
	if err != nil {
		buffer.Push()
		return String{}, err
	}

	if startQuoteToken.Type != lex.QUOTE {
		buffer.Push()
		return String{}, parse.NoGrammarParse
	}

	//capture everything until the next quote that isn't escaped (\)
	var currentToken lex.Token
	var previousToken lex.Token
	var stringBuilder strings.Builder

	for {
		//TODO: support the ESC fragment
		previousToken = currentToken
		currentToken, err = buffer.Pop()
		if err != nil {
			buffer.Push()
			buffer.PushTimes(stringBuilder.Len())
			return String{}, err
		}

		if currentToken.Type == lex.QUOTE && previousToken.Type != lex.BACK_SLASH {
			//a quote that isn't escaped
			break
		} else if currentToken.Type == lex.BACK_SLASH {
			//don't write backslashes to the string; we'll handle what we write on the next token
			continue
		}

		stringBuilder.WriteRune(currentToken.Literal)
	}

	return String{Value: stringBuilder.String()}, nil
}
