package lex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var periodToken = Token{
	Type:    PERIOD,
	Literal: '.',
}

var spaceToken = Token{
	Type:    WHITE_SPACE,
	Literal: ' ',
}

var parenthesisStartToken = Token{
	Type:    PARENTHESIS_START,
	Literal: '(',
}

var parenthesisEndToken = Token{
	Type:    PARENTHESIS_END,
	Literal: ')',
}

var atToken = Token{
	Type:    AT_SIGN,
	Literal: '@',
}

var quoteToken = Token{
	Type:    QUOTE,
	Literal: '\'',
}

var plusToken = Token{
	Type:    PLUS,
	Literal: '+',
}

var dashToken = Token{
	Type:    DASH,
	Literal: '-',
}

var slashToken = Token{
	Type:    SLASH,
	Literal: '/',
}

var starToken = Token{
	Type:    STAR,
	Literal: '*',
}

var backSlashToken = Token{
	Type:    BACK_SLASH,
	Literal: '\\',
}

func TestLexerWithFhirPath(t *testing.T) {
	expectedTokens := []Token{{
		Type:    ALPHA,
		Literal: 'e',
	}, periodToken, {
		Type:    ALPHA,
		Literal: 'r',
	}, periodToken, {
		Type:    ALPHA,
		Literal: 'o',
	}, parenthesisStartToken, {
		Type:    ALPHA,
		Literal: 'P',
	}, spaceToken, {
		Type:    ALPHA,
		Literal: 'o',
	}, spaceToken, {
		Type:    ALPHA,
		Literal: 'S',
	}, parenthesisEndToken, periodToken, {
		Type:    ALPHA,
		Literal: 'i',
	}}

	tokens, err := NewLexer("e.r.o(P o S).i").Lex()

	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens)
}

func TestRandomTokens(t *testing.T) {
	expectedTokens := []Token{parenthesisEndToken, parenthesisStartToken, periodToken, atToken, {
		Type:    NUMERIC,
		Literal: '2',
	}, backSlashToken, quoteToken, {
		Type:    ALPHA,
		Literal: 'o',
	}, plusToken, dashToken, {
		Type:    WHITE_SPACE,
		Literal: '\n',
	}, slashToken, starToken}

	tokens, err := NewLexer(")(.@2\\'o+-\n/*").Lex()

	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens)
}
