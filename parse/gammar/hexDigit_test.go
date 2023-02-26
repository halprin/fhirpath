package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestHexDigitDoesntParseImmediatelyDueToEof(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer(""),
	}

	_, err := NewHexDigit(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)
}

func TestHexDigitDoesntParseImmediatelyDueToNoNumericOrAlpha(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("+"),
	}

	_, err := NewHexDigit(tokenBuffer)

	assert.ErrorIs(t, err, parse.NoGrammarParse)
}

func TestHexDigitDoesntParseImmediatelyDueToBadLowercaseAlpha(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("g"),
	}

	_, err := NewHexDigit(tokenBuffer)

	assert.ErrorIs(t, err, parse.NoGrammarParse)
}

func TestHexDigitDoesntParseImmediatelyDueToBadUppercaseAlpha(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("G"),
	}

	_, err := NewHexDigit(tokenBuffer)

	assert.ErrorIs(t, err, parse.NoGrammarParse)
}

func TestHexDigitParsesWithNumber(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("2"),
	}

	hexDigit, err := NewHexDigit(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, '2', hexDigit.Digit)
}

func TestHexDigitParsesWithLowercaseAlpha(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("f"),
	}

	hexDigit, err := NewHexDigit(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 'f', hexDigit.Digit)
}

func TestHexDigitParsesWithUppercaseAlpha(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("F"),
	}

	hexDigit, err := NewHexDigit(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 'F', hexDigit.Digit)
}
