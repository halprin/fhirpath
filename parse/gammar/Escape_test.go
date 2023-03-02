package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestEscapeDoesntParseImmediatelyDueToEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer(""),
	}

	_, err := NewEscape(tokenBuffer)

	assert.ErrorIs(t, err, io.EOF)
}

func TestEscapeWorks(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("\\t"),
	}

	escape, err := NewEscape(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, '\t', escape.Character)
}

func TestEscapeWorksWithUnicode(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("\\u304C"),
	}

	escape, err := NewEscape(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, 'が', escape.Character)
}

func TestEscapeFailsWithUnicode(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("\\u304g"),
	}

	_, err := NewEscape(tokenBuffer)
	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.BACK_SLASH,
		Literal: '\\',
	}, nextToken)
}

func TestEscapeFailsWithNoBackslash(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("g"),
	}

	_, err := NewEscape(tokenBuffer)
	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: 'g',
	}, nextToken)
}

func TestEscapeFailsWithBadBackslash(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("\\g"),
	}

	_, err := NewEscape(tokenBuffer)
	assert.ErrorIs(t, err, parse.NoGrammarParse)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.BACK_SLASH,
		Literal: '\\',
	}, nextToken)
}

func TestEscapeFailsWithEarlyEof(t *testing.T) {
	tokenBuffer := &parse.TokenBuffer{
		Lexer: lex.NewLexer("\\"),
	}

	_, err := NewEscape(tokenBuffer)
	assert.ErrorIs(t, err, io.EOF)

	nextToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, lex.Token{
		Type:    lex.BACK_SLASH,
		Literal: '\\',
	}, nextToken)
}
