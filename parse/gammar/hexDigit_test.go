package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexDigitWorks(t *testing.T) {
	tokenBuffer := parse.NewTokenBuffer(lex.NewLexer("8"))

	hexDigitGrammar, err := NewHexDigit(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, '8', hexDigitGrammar.Value)
}

func TestHexDigitFailsForNonHex(t *testing.T) {
	tokenBuffer := parse.NewTokenBuffer(lex.NewLexer("g"))

	_, err := NewHexDigit(tokenBuffer)
	assert.ErrorIs(t, err, parse.NoGrammarParse)
}

func TestHexDigitSplitsToken(t *testing.T) {
	tokenBuffer := parse.NewTokenBuffer(lex.NewLexer("82345"))

	hexDigitGrammar, err := NewHexDigit(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, '8', hexDigitGrammar.Value)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.NUMBER,
		Literal: "2345",
	}, nextToken)
}
