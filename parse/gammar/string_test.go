package gammar

import (
	"github.com/halprin/fhirpath/lex"
	"github.com/halprin/fhirpath/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringWorks(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("'keep a true statement or no +'and"),
	}

	stringGrammar, err := NewString(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, "keep a true statement or no +", stringGrammar.Value)
}

func TestStringEmpty(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("''and"),
	}

	stringGrammar, err := NewString(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, "", stringGrammar.Value)
}

func TestStringWithQuoteInIt(t *testing.T) {
	tokenBuffer := parse.TokenBuffer{
		Lexer: lex.NewLexer("'I like \\' in my strings'"),
	}

	stringGrammar, err := NewString(tokenBuffer)
	assert.NoError(t, err)

	assert.Equal(t, "I like ' in my strings", stringGrammar.Value)
}

//TODO: support needs to be added for escapes
//func TestStringWithATabEscape(t *testing.T) {
//	tokenBuffer := parse.TokenBuffer{
//		Lexer: lex.NewLexer("'we have a \\tab in it'"),
//	}
//
//	stringGrammar, err := NewString(tokenBuffer)
//	assert.NoError(t, err)
//
//	assert.Equal(t, "we have a \tab in it", stringGrammar.Value)
//}
