package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestUnicodeDoesntParseImmediatelyDueToEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer(""),
	}

	_, err := NewUnicode(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)
}

func TestUnicodeWorks(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("u0061"),
	}

	unicode, err := NewUnicode(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 'a', unicode.Character)
}

func TestUnicodeWorksForUnicode(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("u304C"),
	}

	unicode, err := NewUnicode(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 'が', unicode.Character)
}

func TestUnicodeFailsForNoHex(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("u006g"),
	}

	_, err := NewUnicode(tokenBuffer)
	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: 'u',
	}, nextToken)
}

func TestUnicodeFailsForEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("u006"),
	}

	_, err := NewUnicode(tokenBuffer)
	assert.ErrorIs(t, err, io.EOF)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: 'u',
	}, nextToken)
}
