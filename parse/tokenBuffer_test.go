package parse

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPopBringsTheTokensInTheCorrectOrder(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("dogcow.moof"),
	}

	firstToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)
	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: "dogcow",
	}, firstToken)

	secondToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)
	assert.Equal(t, lex.Token{
		Type:    lex.PERIOD,
		Literal: ".",
	}, secondToken)

	thirdToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)
	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA,
		Literal: "moof",
	}, thirdToken)

	_, err = tokenBuffer.Pop()
	assert.Error(t, err)
}

func TestPopAndPushBringsBackTheSameToken(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("dogcow.moof"),
	}

	firstToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	tokenBuffer.Push()

	tokenAgain, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, firstToken, tokenAgain)
}

func TestMultiplePopAndPushBringsBackTheSameToken(t *testing.T) {
	tokenBuffer := TokenBuffer{
		Lexer: lex.NewLexer("dogcow.moof.true"),
	}

	firstToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)
	_, err = tokenBuffer.Pop()
	assert.NoError(t, err)
	_, err = tokenBuffer.Pop()
	assert.NoError(t, err)
	_, err = tokenBuffer.Pop()
	assert.NoError(t, err)

	tokenBuffer.Push()
	tokenBuffer.Push()
	tokenBuffer.Push()
	tokenBuffer.Push()

	tokenAgain, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, firstToken, tokenAgain)
}
