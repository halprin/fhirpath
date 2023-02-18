package parse

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNumberDoesntParseImmediatelyDueToEof(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer(""),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)
}

func TestNumberDoesntParseImmediatelyDueToNoNumber(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("dogcow.moof"),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, NoGrammarParse)
}

func TestNumberParsesInteger(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("26abc"),
	}

	number, err := NewNumber(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 26, number.ValueInt)
}

func TestNumberFailsParseAfterPeriod(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("26.abc"),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, NoGrammarParse)
}

func TestNumberParsesFloat(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("26.32abc"),
	}

	number, err := NewNumber(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 26.32, number.ValueFloat)
}
