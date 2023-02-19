package parse

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPopBringsTheTokensInTheCorrectOrder(t *testing.T) {
	tokenBuffer := NewTokenBuffer(lex.NewLexer("dogcow.moof"))

	firstToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)
	assert.Equal(t, lex.Token{
		Type:    lex.ALPHA_NUMERIC,
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
		Type:    lex.ALPHA_NUMERIC,
		Literal: "moof",
	}, thirdToken)

	_, err = tokenBuffer.Pop()
	assert.Error(t, err)
}

func TestPopAndPushBringsBackTheSameToken(t *testing.T) {
	tokenBuffer := NewTokenBuffer(lex.NewLexer("dogcow.moof"))

	firstToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	tokenBuffer.Push()

	tokenAgain, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, firstToken, tokenAgain)
}

func TestMultiplePopAndPushBringsBackTheSameToken(t *testing.T) {
	tokenBuffer := NewTokenBuffer(lex.NewLexer("dogcow.moof.true"))

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

func TestInsertingATokenWithPopingAndPushing(t *testing.T) {
	tokenBuffer := NewTokenBuffer(lex.NewLexer("dogcow.moof.true"))

	_, _ = tokenBuffer.Pop()
	periodToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	tokenBuffer.Push()

	insertedToken := lex.Token{
		Type:    lex.SLASH,
		Literal: "/",
	}

	tokenBuffer.PushToken(insertedToken)

	actualInsertedToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, insertedToken, actualInsertedToken)

	afterInsertedToken, err := tokenBuffer.Pop()
	assert.NoError(t, err)

	assert.Equal(t, periodToken, afterInsertedToken)
}
