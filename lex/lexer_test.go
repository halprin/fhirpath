package lex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var periodToken = Token{
	Type:    PERIOD,
	Literal: ".",
}

var spaceToken = Token{
	Type:    WHITE_SPACE,
	Literal: " ",
}

var parenthesisStartToken = Token{
	Type:    PARENTHESIS_START,
	Literal: "(",
}

var parenthesisEndToken = Token{
	Type:    PARENTHESIS_END,
	Literal: ")",
}

var trueToken = Token{
	Type:    TRUE,
	Literal: "true",
}

var falseToken = Token{
	Type:    FALSE,
	Literal: "false",
}

var andToken = Token{
	Type:    AND,
	Literal: "and",
}

var orToken = Token{
	Type:    OR,
	Literal: "or",
}

var atToken = Token{
	Type:    AT_SIGN,
	Literal: "@",
}

var quoteToken = Token{
	Type:    QUOTE,
	Literal: "'",
}

var plusToken = Token{
	Type:    PLUS,
	Literal: "+",
}

var dashToken = Token{
	Type:    DASH,
	Literal: "-",
}

var slashToken = Token{
	Type:    SLASH,
	Literal: "/",
}

var starToken = Token{
	Type:    STAR,
	Literal: "*",
}

var backSlashToken = Token{
	Type:    BACK_SLASH,
	Literal: "\\",
}

func TestLexerWithFhirPath(t *testing.T) {
	expectedTokens := []Token{{
		Type:    ALPHA,
		Literal: "entry",
	}, periodToken, {
		Type:    ALPHA,
		Literal: "resource",
	}, periodToken, {
		Type:    ALPHA,
		Literal: "ofType",
	}, parenthesisStartToken, {
		Type:    ALPHA,
		Literal: "Patient",
	}, spaceToken, {
		Type:    OR,
		Literal: "or",
	}, spaceToken, {
		Type:    ALPHA,
		Literal: "ServiceRequest",
	}, parenthesisEndToken, periodToken, {
		Type:    ALPHA,
		Literal: "id",
	}}

	tokens, err := NewLexer("entry.resource.ofType(Patient or ServiceRequest).id").Lex()

	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens)
}

func TestRandomTokens(t *testing.T) {
	expectedTokens := []Token{parenthesisEndToken, parenthesisStartToken, trueToken, periodToken, falseToken, atToken, {
		Type:    NUMERIC,
		Literal: "26",
	}, backSlashToken, andToken, quoteToken, orToken, plusToken, dashToken, {
		Type:    WHITE_SPACE,
		Literal: " \t\n\r",
	}, slashToken, starToken}

	tokens, err := NewLexer(")(true.false@26\\and'or+- \t\n\r/*").Lex()

	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens)
}
