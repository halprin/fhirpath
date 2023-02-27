package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNumberDoesntParseImmediatelyDueToEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer(""),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)
}

func TestNumberDoesntParseImmediatelyDueToNoNumber(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("dogcow.moof"),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: 'd',
	}, nextToken)
}

func TestNumberParsesIntegerWithEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("26"),
	}

	number, err := NewNumber(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 26, number.ValueInt)
}

func TestNumberParsesIntegerWithMoreTokens(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("26abc"),
	}

	number, err := NewNumber(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 26, number.ValueInt)
}

func TestNumberFailsParseAfterPeriod(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("26.abc"),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.NUMERIC,
		Literal: '2',
	}, nextToken)
}

func TestNumberFailsParseAfterPeriodWithEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("26."),
	}

	_, err := NewNumber(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.NUMERIC,
		Literal: '2',
	}, nextToken)
}

func TestNumberParsesFloat(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("26.32"),
	}

	number, err := NewNumber(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 26.32, number.ValueFloat)
}
